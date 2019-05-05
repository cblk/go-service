FROM registry.cn-shanghai.aliyuncs.com/ybase/alpine:latest

RUN rm -rf /app
COPY main /app/go-service
WORKDIR /app

ENTRYPOINT ["./go-service","server"]