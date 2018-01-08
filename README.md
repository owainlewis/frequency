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
      - FOO=bar
    steps:
      - echo "Running tests..."
      - go build main.go
      - mv main {{ .Values.OutputDirectory }}
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
   "commands":[
      "go build -v main.go",
      "mv ./main $OUTPUT_DIR",
      "ls $OUTPUT_DIR"
   ]
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

The smallest unit of execution in KCD. A stage is the execution of some commands inside a base image.

#### Context

A job is executed within a context. This means that a job executes with some metadata about the branch etc

### Build

A build has a message, commit and hash and checks out code from VCS and then executes a job based on the contents of a config.yml file.

### Workflow

A workflow describes a series of jobs to be run in a specific order.

## TODO

* Get log streaming from container to upstream API to capture output.
* Parser
* UI basics done
