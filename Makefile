.PHONY: build
build:
	@go build cmd/main.go

.PHONY: clean
clean:
	@rm -rf bin

.PHONY: test
test:
	@go test ./...

.PHONY: run
run:
	@go run cmd/main.go \
	--kubeconfig=${HOME}/.kube/config \
	--v=4 \
	--logtostderr=true
