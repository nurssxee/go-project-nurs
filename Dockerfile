FROM golang:1.24.3-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o main cmd/main.go

FROM alpine:latest

COPY --from=builder /app/main /root/
COPY ../internal/db/migrations ./internal/db/migrations

CMD ["/root/main"]
