IMAGE = frequency-agent
VERSION = $(shell git describe --always)

.PHONY: build
build:
	docker build -t $(IMAGE):$(VERSION)

.PHONY: tag
tag:
	docker tag $(IMAGE):$(VERSION) owainlewis/$(IMAGE):$(VERSION)
