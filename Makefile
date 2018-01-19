.PHONY: build
build:
	@go build cmd/main.go

.PHONY: clean
clean:
	@rm -rf bin

.PHONY: test
test:
	@go test ./...

.PHONY: job
job:
	@curl -iX POST localhost:3000/api/v1/jobs --data-binary "@examples/job.json"

.PHONY: run
run:
	@go run cmd/main.go \
	--kubeconfig=${HOME}/.kube/config \
	--v=4 \
	--logtostderr=true
