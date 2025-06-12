FROM golang:1.24-alpine as builder

ENV GO111MODULE=on \
      CGO_ENABLED=0 \
      GOOS=linux \
      GOARCH=amd64

WORKDIR /app

COPY . .

WORKDIR /app/transaction-service

RUN go build -o transaction-service cmd/main.go

EXPOSE 50053

CMD ["./transaction-service"]
