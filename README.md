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

A task runs code inside a Docker container. 

Optionally you may clone source code prior to running the task for CI like problems.

```yaml
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

```javascript
var frequency = require('frequency');

var taskDef = {image: golang, run: {command: "ci.sh"}};
var task = frequency.NewTask(taskDef);

task.submit();
```

## Misc

### Job

## Trigger (cron, webhook etc)

A trigger is used to execute a job when something happens. This could be a cron schedule or an external event.

```javascript

var trigger = frequency.NewCronTrigger(task: ping);

trigger.create();
```

