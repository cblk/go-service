FROM registry.cn-shanghai.aliyuncs.com/ybase/alpine:latest

RUN rm -rf /app
COPY main /app/main
COPY config/config.yml /app/config/config.yml
WORKDIR /app

ENTRYPOINT ["./main"]
