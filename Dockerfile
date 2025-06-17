# Build stage
FROM golang:1.24 as builder

WORKDIR /app
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o url-shortener

# Runtime stage
FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/url-shortener .

EXPOSE 8080
CMD ["./url-shortener"]