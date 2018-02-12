# Frequency.io

Frequency is a CI/CD system designed to run inside Kubernetes inspired by
CircleCI and Wercker but open source and running atop Kubernetes.

Kubernetes offers many advantages for CI/CD such as:

#### Convenience

* Run your CI jobs on any Kubernetes cluster
* Build jobs get executed inside Kubernetes pods for infinite elasticity and scale

#### Security

* Run CI on your own private infrastructure
* Utilise all the native Kubernetes tooling and security.

## Components

#### API

The API server stores and triggers tasks to be executed.

### Tasks
TODO

### Builds
TODO

### Projects
TODO

### Triggers
TODO

#### CLI

The CLI lets you execute CI actions locally.

```yaml
# Simple task that checks out code and runs unit tests
image: golang
workspace: /go/src/github.com/wercker/getting-started-golang
checkout:
  url: https://github.com/wercker/getting-started-golang.git
steps:
  - go test ./...
  - echo $(env)
```

Complex example

```yaml
kind: kubernetes.io/Pod
image: golang
workspace: /go/src/github.com/wercker/getting-started-golang
env:
  - name: DATABASE_USERNAME
    valueFrom:
      secretKeyRef:
        name: mysql-credentials
        key: username
  - name: DATABASE_PASSWORD
    valueFrom:
      secretKeyRef:
        name: mysql-credentials
        key: password
checkout:
  url: https://github.com/wercker/getting-started-golang.git
run:
  command:
    - bash
    - "-exc"
  args:
    - go test ./...; echo $(pwd); ls -la; echo $(env)
```

Run the task against the API server and it will execute a CI execution of the unit tests for this project.

```
fq task create -f examples/yaml/task.yaml
```

#### Triggers

TODO explain

#### Events

TODO explain

#### Orchestrator (Choreograph)

This engine is responsible for orchestrating piplines of tasks to be executed.

