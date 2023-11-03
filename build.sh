#!/usr/bin/env sh

set -xe
trap 'exit 1' INT

go run ./tools/writetlds/
go build -o ./tldeets ./cmd/tldeets/main.go
