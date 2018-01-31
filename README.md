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

## Task

```
image: golang
workspace: /
run:
  command:
    - bash
    - -exc
  args:
  - |
    whoami
    env
    go version
```

## Sample build manifest

Build manifests are kept inside your project source. The following example shows a build
manifest for building a golang binary inside a container.

Notice that environment variables are stored as Kubernetes secrets.

```yaml
version: 1
tasks:
  wait-60-seconds:
    kind: Wait
    duration: 60
  build:
    kind: CI
    spec:
      image: golang
      workspace: /go/src/github.com/owainlewis/kcd
      environment:
        values:
          GOOS: linux
          GOARCH: amd64
        secretRef:
          name: my-build-secrets
      command:
        cmd: ./ci/build.sh
        args: []
```

## Getting Started

TODO
