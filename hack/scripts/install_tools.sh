#!/usr/bin/env bash

echo 'Installing go tools'

export GO111MODULE=off

go get -u golang.org/x/tools/cmd/goimports
go get -u github.com/onsi/ginkgo/ginkgo
go get -u github.com/vektra/mockery
go get github.com/golangci/golangci-lint/cmd/golangci-lint
go get -u golang.org/x/tools/cmd/cover
go get -u github.com/mattn/goveralls
