gen:
	make -C ../../protobuf/
	@go install ./...

move: gen
	sudo mv ${GOPATH}/bin/protoc-gen-orm /usr/local/bin/

all: move gen
