FROM golang:1.21-alpine AS builder
WORKDIR /app

COPY clashbot/go.mod clashbot/go.sum ./
RUN go mod download

COPY clashbot/ .

RUN CGO_ENABLED=0 GOOS=linux go build -o main .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY --from=builder /app/main .
COPY --from=builder /app/config.json .

CMD ["./main"]
