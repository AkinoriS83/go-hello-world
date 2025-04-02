# Build stage
FROM golang:1.24 AS builder

WORKDIR /app
COPY . .

# Dependency resolution & build
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -o myapp

# Lightweight stage for execution
FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/myapp /app/myapp

EXPOSE 8080

ENTRYPOINT  ["./myapp"]
