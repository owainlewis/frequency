# KCD

A Kubernetes native CI/CD system. Inspired by CircleCI and Wercker but open source and running atop Kubernetes.

## Concepts + Design

A simple API that can be used to launch Kubernetes Jobs.

A simple YAML format for builds

```yaml
version: 1
jobs:
  build:
    image: golang
    workspace: /go/src/github.com/owainlewis/kcd
    environment:
      DOCKER_USERNAME: xxx
      DOCKER_PASSWORD: xxx
    command: 
      cmd: ./ci/build.sh
      args: []
```

## API

### Jobs

Invoke a job execution

`POST /api/v1/jobs`

```json
{
   "name":"hello-kcd",
   "workspace":"/workspace",
   "image":"golang",
   "command": {
      "cmd": "echo",
      "args": ["Hello, World!"]
   }
}
```

```
➜  ~ curl -iX POST localhost:3000/api/v1/jobs -d '{"name":"api-example","workspace":"/workspace","image":"golang","commands":["echo \"BOOM\""]}'
HTTP/1.1 202 Accepted
Date: Sun, 07 Jan 2018 15:40:28 GMT
Content-Length: 108
Content-Type: text/plain; charset=utf-8

{"name":"api-example","workspace":"/workspace","image":"golang","commands":["echo \"BOOM\""],"source":null}
➜  ~ kubectl logs kcd-x00zv
BOOM
```

## Domain Language

Terminology reference.

### Project

A project is typically a software project that lives in version control and contains a kcd.yml file.

### Job

A job is a template describing the execution steps inside in KCD.

### Build

A build is an invokation of a job.

### Workflow

A workflow describes a series of jobs to be run in a specific order.

## TODO

* Get log streaming from container to upstream API to capture output.
* Parser
* UI basics done
