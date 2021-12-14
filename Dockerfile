FROM golang:1.17-alpine3.14 AS builder
RUN apk --no-cache add tzdata
WORKDIR /app
COPY . .
RUN go env -w GOPROXY=https://goproxy.cn && CGO_ENABLED=0 go build -ldflags "-w -s" -o main

FROM alpine:3.14
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
ENV TZ=Asia/Shanghai
COPY run.sh .
COPY static ./static

CMD ["sh", "./run.sh"]
