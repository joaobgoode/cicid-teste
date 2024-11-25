FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.* ./
RUN go mod tidy

COPY . .

RUN go build -o /app/bin/main ./cmd

FROM alpine:latest

RUN apk add --no-cache ca-certificates

WORKDIR /app

COPY --from=builder /app/bin/main .

EXPOSE 8080

CMD ["./main"]

