# Variables
PROJECT = ormpb
REPO_ROOT = ${GOPATH}/src/github.com/gophersbd
ROOT = ${REPO_ROOT}/${PROJECT}

regenerate:
	@echo WARNING! THIS RULE IS PROBABLY NOT RIGHT FOR YOUR INSTALLATION
	@prototool gen option.proto
	protoc --go_out=../. -I${ROOT} ${ROOT}/protobuf/option.proto

