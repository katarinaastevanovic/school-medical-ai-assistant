FROM golang:1.26-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o medical-service .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/medical-service .
EXPOSE 8083
CMD ["./medical-service"]