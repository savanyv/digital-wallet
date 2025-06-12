FROM golang:1.24-alpine as builder

ENV GO111MODULE=on \
      CGO_ENABLED=0 \
      GOOS=linux \
      GOARCH=amd64

WORKDIR /app

COPY . .

WORKDIR /app/auth-service

RUN go build -o auth-service cmd/main.go

EXPOSE 50050

CMD ["./auth-service"]
