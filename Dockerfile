FROM golang:1.16.5-alpine3.14 AS builder
WORKDIR /app
COPY . .
RUN go env -w GOPROXY=https://goproxy.cn && CGO_ENABLED=0 go build -ldflags "-w -s" -o main

FROM alpine:3.14.0
WORKDIR /app
COPY --from=builder /app/main .
COPY run.sh .
COPY static ./static

CMD ["sh", "./run.sh"]
