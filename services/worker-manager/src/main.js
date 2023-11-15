import '../../prelude.js';
import loader from 'taskcluster-lib-loader';
import taskcluster from 'taskcluster-client';
import { App } from 'taskcluster-lib-app';
import { MonitorManager } from 'taskcluster-lib-monitor';
import config from 'taskcluster-lib-config';
import SchemaSet from 'taskcluster-lib-validate';
import libReferences from 'taskcluster-lib-references';
import exchanges from './exchanges.js';
import builder from './api.js';
import { Estimator } from './estimator.js';
import { Client, pulseCredentials } from 'taskcluster-lib-pulse';
import tcdb from 'taskcluster-db';
import { Provisioner } from './provisioner.js';
import { Providers } from './providers/index.js';
import { WorkerScanner } from './worker-scanner.js';
import { WorkerPool, WorkerPoolError, Worker } from './data.js';
import './monitor.js';
import { fileURLToPath } from 'url';

let load = loader({
  cfg: {
    requires: ['profile'],
    setup: ({ profile }) => config({
      profile,
      serviceName: 'worker-manager',
    }),
  },

  monitor: {
    requires: ['process', 'profile', 'cfg'],
    setup: ({ process, profile, cfg }) => MonitorManager.setup({
      serviceName: 'worker-manager',
      processName: process,
      verify: profile !== 'production',
      ...cfg.monitoring,
    }),
  },

  db: {
    requires: ["cfg", "process", "monitor"],
    setup: ({ cfg, process, monitor }) => tcdb.setup({
      readDbUrl: cfg.postgres.readDbUrl,
      writeDbUrl: cfg.postgres.writeDbUrl,
      serviceName: 'worker_manager',
      monitor: monitor.childMonitor('db'),
      statementTimeout: process === 'server' ? 30000 : 0,
      dbCryptoKeys: cfg.postgres.dbCryptoKeys,
    }),
  },

  expireWorkerPools: {
    requires: ['cfg', 'monitor', 'db'],
    setup: ({ cfg, monitor, db }, ownName) => {
      return monitor.childMonitor('expireWorkerPools').oneShot(ownName, async () => {
        const expired = await WorkerPool.expire({ db, monitor });
        for (let workerPoolId of expired) {
          monitor.info(`deleted expired worker pool ${workerPoolId}`);
        }
      });
    },
  },

  expireWorkers: {
    requires: ['cfg', 'monitor', 'db'],
    setup: ({ cfg, monitor, db }, ownName) => {
      return monitor.childMonitor('expireWorkers').oneShot(ownName, async () => {
        const count = await Worker.expire({ db, monitor });
        monitor.info(`deleted ${count} workers`);
      });
    },
  },

  expireErrors: {
    requires: ['cfg', 'monitor', 'db'],
    setup: ({ cfg, monitor, db }, ownName) => {
      return monitor.childMonitor('expireErrors').oneShot(ownName, async () => {
        const count = await WorkerPoolError.expire({
          db,
          retentionDays: cfg.app.workerPoolErrorRetentionDays,
        });
        monitor.info(`deleted ${count} expired worker pool errors`);
      });
    },
  },

  schemaset: {
    requires: ['cfg'],
    setup: ({ cfg }) => new SchemaSet({
      serviceName: 'worker-manager',
    }),
  },

  reference: {
    requires: [],
    setup: () => exchanges.reference(),
  },

  generateReferences: {
    requires: ['cfg', 'schemaset'],
    setup: ({ cfg, schemaset }) => libReferences.fromService({
      schemaset,
      references: [builder.reference(), exchanges.reference(), MonitorManager.reference('worker-manager')],
    }).generateReferences(),
  },

  pulseClient: {
    requires: ['cfg', 'monitor'],
    setup: ({ cfg, monitor }) => {
      return new Client({
        namespace: 'taskcluster-worker-manager',
        monitor: monitor.childMonitor('pulse-client'),
        credentials: pulseCredentials(cfg.pulse),
      });
    },
  },

  publisher: {
    requires: ['cfg', 'schemaset', 'pulseClient'],
    setup: async ({ cfg, pulseClient, schemaset }) => await exchanges.publisher({
      rootUrl: cfg.taskcluster.rootUrl,
      schemaset,
      client: pulseClient,
      publish: false,
    }),
  },

  api: {
    requires: [
      'cfg', 'db', 'schemaset', 'monitor', 'providers',
      'publisher', 'notify'],
    setup: async ({
      cfg,
      db,
      schemaset,
      monitor,
      providers,
      publisher,
      notify,
    }) => builder.build({
      rootUrl: cfg.taskcluster.rootUrl,
      context: {
        cfg,
        db,
        monitor,
        providers,
        publisher,
        notify,
      },
      monitor: monitor.childMonitor('api'),
      schemaset,
    }),
  },

  server: {
    requires: ['cfg', 'api'],
    setup: ({ cfg, api }) => App({
      apis: [api],
      ...cfg.server,
    }),
  },

  queue: {
    requires: ['cfg'],
    setup: ({ cfg }) => new taskcluster.Queue(cfg.taskcluster),
  },

  notify: {
    requires: ['cfg'],
    setup: ({ cfg }) => new taskcluster.Notify(cfg.taskcluster),
  },

  estimator: {
    requires: ['cfg', 'queue', 'monitor'],
    setup: ({ cfg, queue, monitor }) => new Estimator({
      queue,
      monitor: monitor.childMonitor('estimator'),
    }),
  },

  providers: {
    requires: ['cfg', 'monitor', 'notify', 'db', 'estimator', 'schemaset'],
    setup: async ({ cfg, monitor, notify, db, estimator, schemaset }) =>
      new Providers().setup({
        cfg, monitor, notify, db, estimator,
        validator: await schemaset.validator(cfg.taskcluster.rootUrl),
      }),
  },

  azureProviderIds: {
    requires: ['providers'],
    setup: async ({ providers }) => {
      const azureIds = [];
      providers.forAll((provider) => {
        if (provider.providerType === 'azure') {
          azureIds.push(provider.providerId);
        }
      });
      return azureIds.join(',');
    },
  },

  workerScanner: {
    requires: ['cfg', 'monitor', 'providers', 'db', 'azureProviderIds'],
    setup: async ({ cfg, monitor, providers, db, azureProviderIds }, ownName) => {
      const workerScanner = new WorkerScanner({
        ownName,
        providers,
        monitor: monitor.childMonitor('worker-scanner'),
        iterateConf: cfg.app.workerScannerIterateConfig || {},
        providersFilter: {
          cond: '<>', // only run for providers that are not Azure
          value: azureProviderIds,
        },
        db,
      });
      await workerScanner.initiate();
      return workerScanner;
    },
  },

  workerScannerAzure: {
    requires: ['cfg', 'monitor', 'providers', 'db', 'azureProviderIds'],
    setup: async ({ cfg, monitor, providers, db, azureProviderIds }, ownName) => {
      const workerScanner = new WorkerScanner({
        ownName,
        providers,
        monitor: monitor.childMonitor('worker-scanner'),
        iterateConf: cfg.app.workerScannerIterateConfig || {},
        providersFilter: {
          cond: '=', // only run for providers that are Azure
          value: azureProviderIds,
        },
        db,
      });
      await workerScanner.initiate();
      return workerScanner;
    },
  },

  provisioner: {
    requires: ['cfg', 'monitor', 'providers', 'notify', 'reference', 'db'],
    setup: async ({ cfg, monitor, providers, notify, reference, db }, ownName) => {
      return new Provisioner({
        ownName,
        monitor: monitor.childMonitor('provisioner'),
        providers,
        notify,
        db,
        reference,
        rootUrl: cfg.taskcluster.rootUrl,
        iterateConf: cfg.app.provisionerIterateConfig || {},
      });
    },
  },

  runProvisioner: {
    requires: ['provisioner'],
    setup: async ({ provisioner }) => await provisioner.initiate(),
  },

}, {
  profile: process.env.NODE_ENV,
  process: process.argv[2],
});

// If this file is executed launch component from first argument
if (process.argv[1] === fileURLToPath(import.meta.url)) {
  load.crashOnError(process.argv[2]);
}

export default load;
