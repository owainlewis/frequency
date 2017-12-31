# KCD

A Kubernetes native CI/CD system. Inspired by CircleCI and Wercker but open source and running atop Kubernetes.

## Concepts + Design

A simple API that can be used to launch Kubernetes Jobs.

A simple YAML format for builds

```yaml
jobs:
  build:
    image: ubuntu:latest
    steps:
    - checkout:
      - git clone https://github.com/owainlewis/kcd.git
    - test:
      - go test
  deploy:
    image: ubuntu:latest
    steps:
    - push:
      - docker login -u$DOCKER_USERNAME -p$DOCKER_PASSWORD
      - docker push myimage:latest wcr.io/oracle/myimage
```
