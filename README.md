## 一、简介

golang实现的web服务模版项目

## 二、依赖

1. 安装docker和docker-compose环境

2.

 ```shell
 go get -u github.com/jstemmer/go-junit-report
 ```

3.

 ```shell
 go get github.com/axw/gocov/...
 go get github.com/AlekSi/gocov-xml
 ```

## 三、使用

1. 编译
    ```shell script
    docker build -t go-service:dev .
    ```
2. 运行
    ```shell script
    docker-compose up -d
    ```