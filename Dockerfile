FROM registry.cn-shanghai.aliyuncs.com/ybase/alpine:latest

RUN rm -rf /app
COPY main /app/portal
COPY kdata /kdata
WORKDIR /app

ENTRYPOINT ["./portal","server"]