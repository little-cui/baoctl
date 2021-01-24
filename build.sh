#!/usr/bin/env bash

script_path=$(cd "$(dirname "$0")"; pwd)

# build
export GOOS=${GOOS:-windows}
export GOARCH=${GOARCH:-386}

go mod download
go mod vendor

go build -o etc/baoctl.exe ${script_path}/cmd/baoctl/main.go

# package
cd ${script_path}/etc
zip -r baoctl.zip .
