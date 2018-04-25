#!/usr/bin/env bash

set -e

OS=$(uname -s)
ARCH=$(uname -m)

PROTOC_VERSION=3.5.1
[[ ${OS} = 'Darwin' ]] && PROTOC_OS="osx" || PROTOC_OS="linux"
PROTOC_ZIP='protoc-'${PROTOC_VERSION}-${PROTOC_OS}-${ARCH}'.zip'

echo 'Install Protoc and Others required library for OS='${OS} 'ARCH='${ARCH}

curl -OL https://github.com/google/protobuf/releases/download/v${PROTOC_VERSION}/${PROTOC_ZIP}
unzip -o ${PROTOC_ZIP} -d protoc3
sudo mv protoc3/bin/protoc /usr/local/bin/
sudo mkdir -p /usr/local/include/google/protobuf
sudo rm -rf /usr/local/include/google/protobuf/*
sudo mv -f protoc3/include/google/protobuf/* /usr/local/include/google/protobuf/
sudo chmod +x /usr/local/bin/protoc
sudo chmod -R 755 /usr/local/include
rm -f ${PROTOC_ZIP}
rm -rf protoc3

go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
go get -u github.com/golang/protobuf/protoc-gen-go

sudo curl -sSL https://github.com/uber/prototool/releases/download/v0.1.0/prototool-$(uname -s)-$(uname -m) \
  -o /usr/local/bin/prototool && \
  sudo chmod +x /usr/local/bin/prototool && \
  prototool -h
