FROM golang:1.17-alpine3.14 AS builder
WORKDIR /app
COPY . .
ARG GOPROXY=https://goproxy.cn
ARG GOPRIVATE=git.dev.yuanben.org
ARG GIT_USER=maowei
ARG GIT_TOKEN=MjgwNTA0ODEwMjE4Ong+3E9Mh09MnD6XTT0hkdI87c+T

RUN sed -i "s/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g" /etc/apk/repositories &&\
    apk add --no-cache git && \
    git config --global url."https://$GIT_USER:$GIT_TOKEN@$GOPRIVATE".insteadOf "https://$GOPRIVATE" && \
    CGO_ENABLED=0 go build -ldflags "-w -s" -o main

FROM alpine:3.14
WORKDIR /app
COPY --from=builder /app/main .
COPY run.sh .
COPY static ./static

CMD ["sh", "./run.sh"]
