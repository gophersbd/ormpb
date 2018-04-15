#!/usr/bin/env bash

echo 'Installing go tools'

go get -u golang.org/x/tools/cmd/goimports
go get -u github.com/Masterminds/glide
go get -u github.com/sgotti/glide-vc
go get -u github.com/onsi/ginkgo/ginkgo
go get -u github.com/vektra/mockery
go get -u github.com/alecthomas/gometalinter
gometalinter --install
