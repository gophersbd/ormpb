# Makefile includes some useful commands to build or format incentives
# More commands could be added

# Variables
PROJECT = ormpb
REPO_ROOT = ${GOPATH}/src/github.com/gophersbd
ROOT = ${REPO_ROOT}/${PROJECT}

LINTER_PKGS = $(shell glide nv)
LINTER_EXCLUDE = "(^|/)mocks/|(^|/)mock_.*\.go|(^|/)(_)?tests/|(^|/)vendor/|(^|/)example/"

PKGS := $(shell go list ./... | grep -v /vendor | grep -v /tests)

fmt: gen
	@goimports -w *.go cmd pkg tests
	@gofmt -s -w *.go cmd pkg tests
	@prototool format -w protobuf/

compile: fmt
	@prototool compile protobuf/
	@go install . ./cmd/...
	@prototool compile examples/

dep:
	glide up -v
	glide vc --only-code --no-tests

build: compile
	CGO_ENABLED=0 GOOS=linux go build -a -tags netgo -ldflags '-w' -o bin/$(PROJECT) ./cmd/protoc-gen-orm

gen:
	make -C protobuf/

check:
	@prototool format -l protobuf/
	@prototool lint protobuf/
	@gometalinter                        \
         --exclude=${LINTER_EXCLUDE}     \
         --disable-all                   \
         --enable=vet                    \
         --enable=vetshadow              \
         --enable=deadcode               \
         --enable=golint                 \
         --enable=varcheck               \
         --enable=errcheck               \
         --enable=ineffassign            \
         --enable=unconvert              \
         --enable=goconst                \
         --enable=goimports              \
         --enable=misspell               \
         --min-occurrences=5             \
         --enable=gofmt                  \
         --deadline=1000s                \
         ${LINTER_PKGS}

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

test:
	@prototool compile examples
	@go test -timeout $(TEST_TIMEOUT)s $(TEST_ARGS) $(TEST_PKGS)

test-e2e:
	@go test ./tests/e2e/...

clean:
	@rm -rf bin
	@go clean

tools:
	@go get -u golang.org/x/tools/cmd/goimports
	@go get -u github.com/Masterminds/glide
	@go get -u github.com/sgotti/glide-vc
	@go get -u github.com/onsi/ginkgo/ginkgo
	@go get -u github.com/vektra/mockery
	@go get -u golang.org/x/tools/cmd/cover
	@go get -u github.com/mattn/goveralls

cover:
	@go test -v -covermode=count -coverprofile=coverage.out $(TEST_PKGS)
	@$(GOPATH)/bin/goveralls -coverprofile=coverage.out -service=travis-ci -repotoken $(COVERALLS_TOKEN)