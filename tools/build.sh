#!/bin/sh -eux

if [ "$(uname)" = 'Darwin' ]; then
    go build -o genkey cmd/genkey/main.go
    go build -o attach cmd/attach/main.go
    go build -o detach cmd/detach/main.go
    go build -o manhours cmd/manhours/main.go
fi
