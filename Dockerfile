FROM registry.cn-shanghai.aliyuncs.com/ybase/golang-cgo:1.17-alpine3.14 AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 go build -ldflags "-w -s" -o main

FROM alpine:3.14
WORKDIR /app
COPY --from=builder /app/main .
COPY run.sh .
COPY static ./static

CMD ["sh", "./run.sh"]
