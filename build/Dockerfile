FROM golang:alpine3.20 AS BUILD

COPY . /grpc-ping-pong
WORKDIR /grpc-ping-pong
RUN GOOS=linux GOARCH=amd64 go build -o grpc-ping-pong ./cmd/ping_pong/main.go

FROM alpine:3.20

RUN set -eux; \
    addgroup -S -g 1000 pingpong; \
    adduser -S -G pingpong -u 999 pingpong

COPY --from=BUILD --chown=pingpong:pingpong /grpc-ping-pong/grpc-ping-pong /usr/local/bin/

EXPOSE 50051
CMD ["grpc-ping-pong"]
