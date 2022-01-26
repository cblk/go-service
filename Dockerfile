FROM ghcr.io/cblk/golang-cgo:1.17-alpine3.14 AS builder
WORKDIR /app
COPY . .
RUN GIT_SHA=$(git rev-parse --short HEAD) && \
    CGO_ENABLED=0 go build -ldflags "-w -s -X main.sha=${GIT_SHA}" -o main

FROM alpine:3.14
WORKDIR /app
COPY --from=builder /app/main .
COPY run.sh .
COPY static ./static

CMD ["sh", "./run.sh"]
