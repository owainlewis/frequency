# API

## Jobs

Invoke a new job. 

When invoked, this endpoint will create a new Kubernetes pod and 
execute the steps provided. If a source option is given, a sidecar container
will clone the project source code and mount it into the workspace.

### POST /api/v1/jobs



