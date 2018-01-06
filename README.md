# KCD

A Kubernetes native CI/CD system. Inspired by CircleCI and Wercker but open source and running atop Kubernetes.

## Concepts + Design

A simple API that can be used to launch Kubernetes Jobs.

A simple YAML format for builds

```yaml
jobs:
  build:
    image: golang
    script:
      - go build cmd/main.go

  test:
    image: golang
    script:
      - go test ./...
```

## Domain Language

Terminology reference.

### Project

A project is typically a software project that lives in version control and contains a kcd.yml file.

### Job

The smallest unit of execution in KCD. A stage is the execution of some commands inside a base image.

#### Context

A job is executed within a context. This means that a job executes with some metadata about the branch etc

### Workflow

A workflow describes a series of jobs to be run in a specific order.

## TODO

* Get log streaming from container to upstream API to capture output.
* Parser
* UI basics done
