FROM golang:1.16-alpine

WORKDIR /app

ENV MODE=development \
    POSTGRES_HOST=localhost \
    POSTGRES_PASSWORD=qwerty \
    POSTGRES_DB=postgres \
    POSTGRES_USER=postgres \
    POSTGRES_PORT=5432 \
    JWT_ACCESS_TOKEN_SECRET=secret \
    JWT_REFRESH_TOKEN_SECRET=secret

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

CMD [ "/app/bin/server" ]
