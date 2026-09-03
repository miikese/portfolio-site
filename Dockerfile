FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY go.mod ./
COPY . .
RUN go build -o portfolio .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/portfolio .
EXPOSE 8080
CMD ["./portfolio"]
