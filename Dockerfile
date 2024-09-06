FROM golang:1.23.0-alpine as builder

COPY . /app/
WORKDIR /app/

RUN go mod download
RUN go build -o ./bin/server cmd/grpc_server/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /app/bin/server .

CMD ["./server"]