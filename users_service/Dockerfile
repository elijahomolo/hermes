FROM golang:1.24-alpine AS builder

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o users_service ./main.go

FROM alpine:3.21 AS final

COPY --from=builder /app/users_service /usr/local/bin/users_service

EXPOSE 8080

CMD ["/usr/local/bin/users_service"]

