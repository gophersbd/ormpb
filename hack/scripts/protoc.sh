#!/usr/bin/env bash

set -e

OS=$(uname -s)
ARCH=$(uname -m)

PROTOC_VERSION=3.5.1
[[ ${OS} = 'Darwin' ]] && PROTOC_OS="osx" || PROTOC_OS="linux"
PROTOC_ZIP='protoc-'${PROTOC_VERSION}-${PROTOC_OS}-${ARCH}'.zip'

echo 'Install Protoc and Others required library for OS='${OS} 'ARCH='${ARCH}

curl -OL https://github.com/google/protobuf/releases/download/v${PROTOC_VERSION}/${PROTOC_ZIP}
sudo unzip -o ${PROTOC_ZIP} -d /usr/local bin/protoc
sudo chmod +x /usr/local/bin/protoc
rm -f ${PROTOC_ZIP}

go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
go get -u github.com/golang/protobuf/protoc-gen-go

sudo curl -sSL https://github.com/uber/prototool/releases/download/v0.1.0/prototool-$(uname -s)-$(uname -m) \
  -o /usr/local/bin/prototool && \
  chmod +x /usr/local/bin/prototool && \
  prototool -h