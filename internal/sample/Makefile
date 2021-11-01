MODULE = $(shell go list -m)
SHELL := /bin/bash
LINT_TOOL=$(shell go env GOPATH)/bin/golangci-lint
GO_PKGS=$(foreach pkg, $(shell go list ./...), $(if $(findstring /vendor/, $(pkg)), , $(pkg)))
GO_FILES=$(shell find . -type f -name '*.go' -not -path './vendor/*')

ENV := local
ifdef $$APP_ENV
ENV := $$APP_ENV
endif

export PROJECT = github.com/arimaulana/point-of-no-return/internal/sample

build:
	env GOOS=linux GOARCH=amd64 go build -o bin/server $(PROJECT)/cmd/server
	env GOOS=linux GOARCH=amd64 go build -o bin/admin $(PROJECT)/cmd/admin
	chmod +x bin/server
	chmod +x bin/admin

build-mac:
	env GOOS=darwin GOARCH=amd64 go build -o bin/server $(PROJECT)/cmd/server
	env GOOS=darwin GOARCH=amd64 go build -o bin/admin $(PROJECT)/cmd/admin
	chmod +x bin/server
	chmod +x bin/admin

run:
	cd cmd/server && go run .

start:
	./bin/server

migrate:
	./bin/admin migrate up

migrate-down:
	./bin/admin migrate down

seed:
	./bin/admin seed test/testdata/seed.sql

test:
	go test ./... -count=1

deps-reset:
	git checkout -- go.mod
	go mod tidy
	go mod vendor

tidy:
	go mod tidy
	go mod vendor

deps-upgrade:
	go get -u -t -d -v ./...

deps-cleancache:
	go clean -modcache

fmt:
	@go fmt $(GO_PKGS)
	@goimports -w -l $(GO_FILES)

$(LINT_TOOL):
	curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b $(shell go env GOPATH)/bin v1.26.0

qc: $(LINT_TOOL)
	$(LINT_TOOL) run --config=.golangci.yaml ./...
	staticcheck ./...