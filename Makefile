.PHONY: docs
docs:
	./scripts/make_swager_docs.sh

.PHONY: push
push:
	./scripts/gitPush.sh

.PHONY: migrate
migrate:
	go run ./cmd/migrate

.PHONY: build
build:
	go build -v ./cmd/api

.PHONY: dev
dev:
	go run ./cmd/api -mode=dev

.PHONY: test
test:
	#go test -v -cover -mode=test -timeout 3s ./...
	#go test -v -race -cover -timeout 3s ./...
	go test -v -cover -timeout 3s ./...

.DEFAULT_GOAL := build