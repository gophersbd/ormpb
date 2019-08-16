# Makefile includes some useful commands to build or format incentives
# More commands could be added

# Variables
PROJECT = ormpb
REPO_ROOT = ${GOPATH}/src/github.com/gophersbd
ROOT = ${REPO_ROOT}/${PROJECT}

PKGS := $(shell go list ./... | grep -v /vendor | grep -v /tests)

fmt: gen
	@goimports -w *.go cmd pkg tests
	@gofmt -s -w *.go cmd pkg tests
	@prototool format -w protobuf/

gen-examples:
	@make -C examples/postgres
	@make -C examples/mysql

compile-examples:
	@prototool compile examples/mysql
	@prototool compile examples/postgres

compile-code: fmt
	@prototool compile protobuf
	@go install . ./cmd/...

compile: fmt compile-code compile-examples

dep:
	go mod download
	go mod vendor
	go mod tidy

build: compile
	CGO_ENABLED=0 GOOS=linux go build -a -tags netgo -ldflags '-w' -o bin/$(PROJECT) ./cmd/protoc-gen-orm

gen:
	@prototool generate protobuf/option.proto

check:
	@prototool format -l protobuf/
	@prototool lint protobuf/
	@golangci-lint run . cmd/... pkg/...

# A user can invoke tests in different ways:
#  - make test runs all tests;
#  - make test TEST_TIMEOUT=10 runs all tests with a timeout of 10 seconds;
#  - make test TEST_PKG=./model/... only runs tests for the model package;
#  - make test TEST_ARGS="-v -short" runs tests with the specified arguments;
#  - make test-race runs tests with race detector enabled.
TEST_TIMEOUT = 60
TEST_PKGS ?= ./cmd/... ./pkg/... .
TEST_TARGETS := test-short test-verbose test-race test-cover
.PHONY: $(TEST_TARGETS) test tests
test-short:   TEST_ARGS=-short
test-verbose: TEST_ARGS=-v
test-race:    TEST_ARGS=-race
test-cover:   TEST_ARGS=-cover
$(TEST_TARGETS): test

test: compile-examples
	@go test -timeout $(TEST_TIMEOUT)s $(TEST_ARGS) $(TEST_PKGS)

test-e2e: gen-examples
	@go test ./tests/e2e/...

clean:
	@rm -rf bin
	@go clean

cover:
	@go test -v -covermode=count -coverprofile=coverage.out $(TEST_PKGS)
	@$(GOPATH)/bin/goveralls -coverprofile=coverage.out -service=travis-ci -repotoken $(COVERALLS_TOKEN)
