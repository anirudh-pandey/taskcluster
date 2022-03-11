# coding=utf-8
#####################################################
# THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT #
#####################################################
# noqa: E128,E201
from ...aio.asyncclient import AsyncBaseClient
from ...aio.asyncclient import createApiClient
from ...aio.asyncclient import config
from ...aio.asyncclient import createTemporaryCredentials
from ...aio.asyncclient import createSession
_defaultConfig = config


class Index(AsyncBaseClient):
    """
    The index service is responsible for indexing tasks. The service ensures that
    tasks can be located by user-defined names.

    As described in the service documentation, tasks are typically indexed via Pulse
    messages, so the most common use of API methods is to read from the index.

    Slashes (`/`) aren't allowed in index paths.
    """

    classOptions = {
    }
    serviceName = 'index'
    apiVersion = 'v1'

    async def ping(self, *args, **kwargs):
        """
        Ping Server

        Respond without doing anything.
        This endpoint is used to check that the service is up.

        This method is ``stable``
        """

        return await self._makeApiCall(self.funcinfo["ping"], *args, **kwargs)

    async def lbheartbeat(self, *args, **kwargs):
        """
        Load Balancer Heartbeat

        Respond without doing anything.
        This endpoint is used to check that the service is up.

        This method is ``stable``
        """

        return await self._makeApiCall(self.funcinfo["lbheartbeat"], *args, **kwargs)

    async def version(self, *args, **kwargs):
        """
        Taskcluster Version

        Respond with the JSON version object.
        https://github.com/mozilla-services/Dockerflow/blob/main/docs/version_object.md

        This method is ``stable``
        """

        return await self._makeApiCall(self.funcinfo["version"], *args, **kwargs)

    async def findTask(self, *args, **kwargs):
        """
        Find Indexed Task

        Find a task by index path, returning the highest-rank task with that path. If no
        task exists for the given path, this API end-point will respond with a 404 status.

        This method is ``stable``
        """

        return await self._makeApiCall(self.funcinfo["findTask"], *args, **kwargs)

    async def listNamespaces(self, *args, **kwargs):
        """
        List Namespaces

        List the namespaces immediately under a given namespace.

        This endpoint
        lists up to 1000 namespaces. If more namespaces are present, a
        `continuationToken` will be returned, which can be given in the next
        request. For the initial request, the payload should be an empty JSON
        object.

        This method is ``stable``
        """

        return await self._makeApiCall(self.funcinfo["listNamespaces"], *args, **kwargs)

    async def listTasks(self, *args, **kwargs):
        """
        List Tasks

        List the tasks immediately under a given namespace.

        This endpoint
        lists up to 1000 tasks. If more tasks are present, a
        `continuationToken` will be returned, which can be given in the next
        request. For the initial request, the payload should be an empty JSON
        object.

        **Remark**, this end-point is designed for humans browsing for tasks, not
        services, as that makes little sense.

        This method is ``stable``
        """

        return await self._makeApiCall(self.funcinfo["listTasks"], *args, **kwargs)

    async def insertTask(self, *args, **kwargs):
        """
        Insert Task into Index

        Insert a task into the index.  If the new rank is less than the existing rank
        at the given index path, the task is not indexed but the response is still 200 OK.

        Please see the introduction above for information
        about indexing successfully completed tasks automatically using custom routes.

        This method is ``stable``
        """

        return await self._makeApiCall(self.funcinfo["insertTask"], *args, **kwargs)

    async def deleteTask(self, *args, **kwargs):
        """
        Remove Task from Index

        Remove a task from the index.  This is intended for administrative use,
        where an index entry is no longer appropriate.  The parent namespace is
        not automatically deleted.  Index entries with lower rank that were
        previously inserted will not re-appear, as they were never stored.

        This method is ``stable``
        """

        return await self._makeApiCall(self.funcinfo["deleteTask"], *args, **kwargs)

    async def findArtifactFromTask(self, *args, **kwargs):
        """
        Get Artifact From Indexed Task

        Find a task by index path and redirect to the artifact on the most recent
        run with the given `name`.

        Note that multiple calls to this endpoint may return artifacts from differen tasks
        if a new task is inserted into the index between calls. Avoid using this method as
        a stable link to multiple, connected files if the index path does not contain a
        unique identifier.  For example, the following two links may return unrelated files:
        * https://tc.example.com/api/index/v1/task/some-app.win64.latest.installer/artifacts/public/installer.exe`
        * https://tc.example.com/api/index/v1/task/some-app.win64.latest.installer/artifacts/public/debug-symbols.zip`

        This problem be remedied by including the revision in the index path or by bundling both
        installer and debug symbols into a single artifact.

        If no task exists for the given index path, this API end-point responds with 404.

        This method is ``stable``
        """

        return await self._makeApiCall(self.funcinfo["findArtifactFromTask"], *args, **kwargs)

    async def heartbeat(self, *args, **kwargs):
        """
        Heartbeat

        Respond with a service heartbeat.

        This endpoint is used to check on backing services this service
        depends on.

        This method is ``stable``
        """

        return await self._makeApiCall(self.funcinfo["heartbeat"], *args, **kwargs)

    funcinfo = {
        "deleteTask": {
            'args': ['namespace'],
            'method': 'delete',
            'name': 'deleteTask',
            'route': '/task/<namespace>',
            'stability': 'stable',
        },
        "findArtifactFromTask": {
            'args': ['indexPath', 'name'],
            'method': 'get',
            'name': 'findArtifactFromTask',
            'route': '/task/<indexPath>/artifacts/<name>',
            'stability': 'stable',
        },
        "findTask": {
            'args': ['indexPath'],
            'method': 'get',
            'name': 'findTask',
            'output': 'v1/indexed-task-response.json#',
            'route': '/task/<indexPath>',
            'stability': 'stable',
        },
        "heartbeat": {
            'args': [],
            'method': 'get',
            'name': 'heartbeat',
            'route': '/__heartbeat__',
            'stability': 'stable',
        },
        "insertTask": {
            'args': ['namespace'],
            'input': 'v1/insert-task-request.json#',
            'method': 'put',
            'name': 'insertTask',
            'output': 'v1/indexed-task-response.json#',
            'route': '/task/<namespace>',
            'stability': 'stable',
        },
        "lbheartbeat": {
            'args': [],
            'method': 'get',
            'name': 'lbheartbeat',
            'route': '/__lbheartbeat__',
            'stability': 'stable',
        },
        "listNamespaces": {
            'args': ['namespace'],
            'method': 'get',
            'name': 'listNamespaces',
            'output': 'v1/list-namespaces-response.json#',
            'query': ['continuationToken', 'limit'],
            'route': '/namespaces/<namespace>',
            'stability': 'stable',
        },
        "listTasks": {
            'args': ['namespace'],
            'method': 'get',
            'name': 'listTasks',
            'output': 'v1/list-tasks-response.json#',
            'query': ['continuationToken', 'limit'],
            'route': '/tasks/<namespace>',
            'stability': 'stable',
        },
        "ping": {
            'args': [],
            'method': 'get',
            'name': 'ping',
            'route': '/ping',
            'stability': 'stable',
        },
        "version": {
            'args': [],
            'method': 'get',
            'name': 'version',
            'route': '/__version__',
            'stability': 'stable',
        },
    }


__all__ = ['createTemporaryCredentials', 'config', '_defaultConfig', 'createApiClient', 'createSession', 'Index']
