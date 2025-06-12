FROM golang:1.24-alpine as builder

ENV GO111MODULE=on \
      CGO_ENABLED=0 \
      GOOS=linux \
      GOARCH=amd64

WORKDIR /app

COPY . .

WORKDIR /app/user-service

RUN go build -o user-service cmd/main.go

EXPOSE 50051

CMD ["./user-service"]
