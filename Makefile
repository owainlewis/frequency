VERSION=$(shell git describe --always)

.PHONY: build
build:
	@go build cmd/main.go

.PHONY: clean
clean:
	@rm -rf bin

.PHONY: test
test:
	@go test ./...

.PHONY: run-task
run-task:
	@curl -iX POST --header "Content-Type: application/json" \
	localhost:3000/api/v1/tasks --data-binary "@examples/task.json"

.PHONY: run-build
run-build:
	@curl -iX POST --header "Content-Type: application/json" \
	localhost:3000/api/v1/projects/1/builds --data-binary "@examples/build.json"


.PHONY: run
run:
	@go run cmd/main.go \
	--kubeconfig=${HOME}/.kube/config \
	--v=4 \
	--logtostderr=true
