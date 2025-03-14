FROM golang:latest AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy
RUN go install github.com/swaggo/swag/cmd/swag@latest

COPY . .
COPY .env /app/.env
COPY .env.local /app/.env.local

RUN make build

FROM debian:bookworm-slim

RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/bin/core-api-template /usr/local/bin/
COPY .env.local /app/.env.local

EXPOSE 8080

CMD ["core-api-template"]
