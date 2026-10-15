FROM golang:1.26-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o chatbot-service .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/chatbot-service .
EXPOSE 8084
CMD ["./chatbot-service"]