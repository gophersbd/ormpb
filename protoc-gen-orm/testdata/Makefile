# Variables
PROJECT = ormpb
REPO_ROOT = ${GOPATH}/src/github.com/gophersbd
ROOT = ${REPO_ROOT}/${PROJECT}

gen:
	@echo WARNING! THIS RULE IS PROBABLY NOT RIGHT FOR YOUR INSTALLATION
	@prototool compile message_options.proto
	@protoc --go_out=../../. -I/usr/local/include -I${ROOT} ${ROOT}/protoc-gen-orm/testdata/message_options.proto
	@protoc --orm_out=../../. -I/usr/local/include -I${ROOT} ${ROOT}/protoc-gen-orm/testdata/message_options.proto
