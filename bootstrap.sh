#!/usr/bin/env bash

set -e

function require {
    command -v $1 2>&1 1>/dev/null || {
        echo "Please install $1 before running this script."
        exit 1
    }
}

set -x

require go
require curl
require glide

go get -u golang.org/x/tools/cmd/goimports
go get -u github.com/lestrrat/go-bindata/...
go get -u git.lukas.moe/sn0w/ropus
