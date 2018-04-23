# Variables
PROJECT = ormpb
REPO_ROOT = ${GOPATH}/src/github.com/gophersbd
ROOT = ${REPO_ROOT}/${PROJECT}

regenerate:
	@echo WARNING! THIS RULE IS PROBABLY NOT RIGHT FOR YOUR INSTALLATION
	@prototool compile message_options.proto
	protoc --orm_out=../. -I${ROOT} ${ROOT}/protoc-gen-orm/testdata/message_options.proto

