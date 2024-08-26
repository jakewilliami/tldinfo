#!/usr/bin/env sh

set -xe
trap 'exit 1' INT

go generate  # runs writetlds for json and go
go build -o ./tldinfo ./cmd/tldinfo/main.go
