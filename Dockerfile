FROM golang:1.16-alpine

WORKDIR /app

COPY ./go.sum ./go.mod ./*.go ./
COPY ./cmd /app/cmd
COPY ./configs /app/configs

RUN go mod download \
    && mkdir bin \
    && go build -o ./bin/server ./cmd/main.go

EXPOSE 8080

CMD [ "./bin/server" ]
