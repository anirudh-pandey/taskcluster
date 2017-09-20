// THIS IS AN AUTOGENERATED FILE. DO NOT EDIT THIS FILE DIRECTLY.

import Client from '../Client';

export default class TreeherderEvents extends Client {
  constructor(options = {}) {
    super({
      baseUrl: '',
      exchangePrefix: 'exchange/taskcluster-treeherder/v1/',
      ...options
    });
  }

  // When a task run is scheduled or resolved, a message is posted to
  // this exchange in a Treeherder consumable format.
  jobs(pattern) {
    const entry = {type:'topic-exchange',exchange:'jobs',name:'jobs',routingKey:[{name:'destination',multipleWords:false,required:true},{name:'project',multipleWords:false,required:true},{name:'reserved',multipleWords:true,required:false}],schema:'http://schemas.taskcluster.net/taskcluster-treeherder/v1/pulse-job.json#'}; // eslint-disable-line

    return this.normalizePattern(entry, pattern);
  }
}
