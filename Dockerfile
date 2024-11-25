FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.* ./

RUN go mod tidy

COPY . .

RUN go build -o main .

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]
