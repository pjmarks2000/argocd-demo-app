FROM golang:1.22 AS builder
WORKDIR /app
COPY app/ .
RUN go build -o server main.go

FROM debian:stable-slim
COPY --from=builder /app/server /usr/local/bin/server
CMD ["server"]
EXPOSE 80
