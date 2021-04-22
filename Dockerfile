FROM golang:alpine AS builder
WORKDIR /app
COPY . .
RUN go env -w GOPROXY=https://goproxy.cn && CGO_ENABLED=0 go build -ldflags "-w -s" -o main

FROM alpine
WORKDIR /app
COPY --from=builder /app/main .
COPY run.sh .
COPY static ./static

CMD ["sh", "./run.sh"]
