#!/usr/bin/env sh

set -xe
trap 'exit 1' INT

go run ./tools/writetlds/ const
go build -o ./tldinfo ./cmd/tldinfo/main.go
