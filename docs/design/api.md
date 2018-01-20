# API

## Jobs

Invoke a new job.

When invoked, this endpoint will create a new Kubernetes pod and
execute the steps provided. If a source option is given, a sidecar container
will clone the project source code and mount it into the workspace.

### POST /api/v1/jobs

Sample payload

```json
{
  "image":"golang",
  "workspace":"/go/src/github.com/owainlewis/frequency-demo-project",
  "env":[
    {
      "name":"SECRET_USERNAME",
      "valueFrom":{
        "secretKeyRef":{
          "name":"mysecret",
          "key":"username"
        }
      }
    }
  ],
  "command":{
    "cmd":"./ci/build.sh",
    "args":[]
  },
  "source":{
    "git_url":"https://github.com/owainlewis/frequency-demo-project.git",
    "git_branch":"master"
  }
}
```



