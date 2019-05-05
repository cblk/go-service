FROM registry.cn-shanghai.aliyuncs.com/ybase/alpine:latest

RUN rm -rf /app
COPY main /app/main
WORKDIR /app

ENTRYPOINT ["./main","server"]