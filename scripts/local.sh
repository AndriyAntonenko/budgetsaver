#!/bin/bash

export MODE=local
export GO111MODULE=on && \
    rm -f ./bin/main && \
    rm -f server.log && \
    go build -o ./bin/main ./cmd/main.go && \
    ./bin/main
