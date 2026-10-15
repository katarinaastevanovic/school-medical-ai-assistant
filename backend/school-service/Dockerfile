FROM golang:1.26-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o school-service .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/school-service .
EXPOSE 8082
CMD ["./school-service"]