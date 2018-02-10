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
--
name: unit-tests
image: golang
workspace: /go/src/github.com/wercker/getting-started-golang
source:
  domain: github.com
  owner: wercker
  repository: getting-started-golang
run:
  command:
    - bash
    - "-exc"
  args:
    - go test ./...
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

