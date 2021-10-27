#!/bin/bash

# Run container
docker run -p 5432:5432 -e POSTGRES_PASSWORD='qwerty' --rm -d --name bs-dev-db postgres

# Run migrations
migrate -path ./schemas -database 'postgres://postgres:qwerty@localhost:5432/postgres?sslmode=disable' up