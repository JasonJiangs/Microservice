package main

import (
	"net/rpc"
	"net/rpc/jsonrpc"
)

// check registered object is valid or not at compile time

// create an interface for method
type HelloServiceInterface interface {
	HelloServer(string, *string) error
}

// need to give parameter when calling the method, 实现了HelloServiceInterface接口
func registerHelloService(svc HelloServiceInterface) {
	rpc.RegisterName("server01", svc)
}

// encap for client

type Client struct {
	c *rpc.Client
}

// 由于使用了c调用call，所以需要初始化c
func InitClient(addr string) Client {
	conn, _ := jsonrpc.Dial("tcp", addr)
	return Client{c: conn}
}

func (this *Client) HelloServer(a string, b *string) error {
	return this.c.Call("server01.HelloServer", a, b)
}
