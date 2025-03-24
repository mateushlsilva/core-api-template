FROM golang:1.23-alpine AS builder

RUN apk update && apk add --no-cache git
RUN go install github.com/swaggo/swag/cmd/swag@latest

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN swag init --parseDependency

COPY src/database/migrations /app/src/database/migrations

RUN go build -o /app/core-api

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/core-api /app/core-api

COPY src/database/migrations /app/src/database/migrations

COPY .env /app/.env
COPY .env.local /app/.env.local

EXPOSE 8080

CMD ["/app/core-api"]
