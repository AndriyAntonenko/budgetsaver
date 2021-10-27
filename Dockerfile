FROM golang:1.16-alpine

WORKDIR /app

# TODO: Multistage build
COPY ./go.sum ./go.mod ./*.go ./
COPY ./pkg /app/pkg
COPY ./cmd /app/cmd
COPY ./scripts/wait-for-it.sh /app/wait-for-it.sh
RUN chmod +x /app/wait-for-it.sh
COPY ./configs /app/configs

RUN apk add --no-cache bash

RUN go mod download \
    && mkdir bin \
    && go build -o ./bin/server ./cmd/main.go

EXPOSE 8080

CMD [ "./bin/server" ]
