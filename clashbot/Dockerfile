FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./main.go
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]

# FROM golang:1.21-alpine AS builder
# WORKDIR /app


# COPY go.mod go.sum ./
# RUN go mod download

# COPY *.go ./
# COPY config/ ./config/
# COPY cr/ ./cr/
# COPY service/ ./service/
# COPY tg/ ./tg/
# COPY config.json ./

# RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# FROM alpine:latest
# RUN apk --no-cache add ca-certificates
# WORKDIR /root/

# COPY --from=builder /app/main .
# COPY --from=builder /app/config.json .

# CMD ["./main"]