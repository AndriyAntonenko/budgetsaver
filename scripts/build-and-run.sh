# /bin/sh

export GO111MODULE=on && \
    rm -f ./bin/main && \
    go build -o ./bin/main ./cmd/main.go && \
    ./bin/main
