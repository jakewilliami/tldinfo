#!/usr/bin/env sh

set -xe
trap 'exit 1' INT

go run ./tools/writetlds/ const
go build -o ./tldeets ./cmd/tldeets/main.go
