FROM golang:1.24-alpine as builder

ENV GO111MODULE=on \
      CGO_ENABLED=0 \
      GOOS=linux \
      GOARCH=amd64

WORKDIR /app

COPY . .

WORKDIR /app/api-gateway

RUN go build -o api-gateway cmd/main.go

EXPOSE 3000

CMD ["./api-gateway"]
