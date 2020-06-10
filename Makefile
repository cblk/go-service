r:
	go run main.go
b:
	go build main.go
s:
	go run main.go server
test:
	go test -v tests/*
docker:
	docker build --file Dockerfile.dev -t registry.cn-shanghai.aliyuncs.com/ybase/go-service:dev .
up:
	docker-compose up -d
restart:
	docker-compose restart api
down:
	docker-compose down