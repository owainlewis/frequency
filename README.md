# KCD

A Kubernetes native CI/CD system. Inspired by CircleCI and Wercker but open source and running atop Kubernetes.

## Concepts + Design

A simple API that can be used to launch Kubernetes Jobs.

A simple YAML format for builds

```yaml
stages:
  build:
    image: golang
    script:
      - go build cmd/main.go
      
  test:
    image: golang
    script:
      - go test ./...
```
