FROM golang:1.24-alpine as builder

ENV GO111MODULE=on \
      CGO_ENABLED=0 \
      GOOS=linux \
       GOARCH=amd64

WORKDIR /app

COPY . .

WORKDIR /app/wallet-service

RUN go build -o wallet-service cmd/main.go

EXPOSE 50052

CMD ["./wallet-service"]
