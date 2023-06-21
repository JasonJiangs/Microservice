# Golang-Microservice

The project implements several small scripts and microservices using Golang and its multiple frameworks.

There are several technologies used:
1. Remote Procedure Call Protocol (RPC)
2. [gRPC](https://grpc.io/)
3. [Protocol Buffers](https://protobuf.dev/)
4. [go-micro](https://github.com/go-micro/go-micro)
5. [Consul](https://www.consul.io/)
5. [Gin](https://gin-gonic.com/)
6. [Redis](https://redis.com/?_ga=2.168453218.258787526.1682134477-94984867.1682134477&_gac=1.47254101.1682134479.CjwKCAjw6IiiBhAOEiwALNqncS8bEal-GSBDztQ_kz6PIeBkgH_iLPlMOWIT2vJFEmxcuD1lLL6p-RoCtsIQAvD_BwE&_gl=1%2A1qo3h22%2A_ga%2AOTQ5ODQ4NjcuMTY4MjEzNDQ3Nw..%2A_ga_8BKGRQKRPV%2AMTY4MjEzNDQ3Ny4xLjEuMTY4MjEzNDQ4Ny41MC4wLjA.)
7. [Alibaba Cloud](https://us.alibabacloud.com/?utm_key=se_1007723047&utm_content=se_1007723047&gad=1&gclid=CjwKCAjw6IiiBhAOEiwALNqncahCRVpweroKxsTkwd-WGdPQb25klMkOBPmueAloKtivb7HbFBw1iRoC98IQAvD_BwE)
8. [GORM](https://gorm.io/)
9. [MySQL](https://www.mysql.com/)
10. [FastDFS](https://github.com/happyfish100/fastdfs)
11. [Nginx](https://www.nginx.com/)

## Practice 1: Chatting Room
### Path: src/Chatting Room
This project implements a chatting room mainly using channel and net library to 
realize the communication between multiple clients and server.

## Practice 2: RPCapi
### Path: src/RPCapi
RPC is a remote procedure call protocol. It is a communication protocol between two computers on a network.
The project implements a simple RPC server and client, and simply encapsulates the function of
server and client.

## Practice 3: gRPC
### Path: src/gRPC
Want to use protobuf, we need to use gRPC as a whole package.

Compile protobuf file:
```Bash
go install github.com/golang/protobuf/protoc-gen-go@latest
protoc --go_out=./ *.proto
# grpc compile instead of protoc
protoc --go_out=plugins=grpc:./ *.proto
```
Install gRPC:
```Bash
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
$ export PATH="$PATH:$(go env GOPATH)/bin"
```

