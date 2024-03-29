package main

import (
	"fmt"
	"net/rpc/jsonrpc"
)

func main01() {
	// dial to connect to the server
	//conn, err := rpc.Dial("tcp", "127.0.0.1:8800")
	conn, err := jsonrpc.Dial("tcp", "127.0.0.1:8800")
	if err != nil {
		fmt.Println("Dial error:", err)
		return
	}
	defer conn.Close()

	// define a variable to receive the response
	var response string
	// call the method of the server
	err = conn.Call("server01.HelloServer", "client01", &response)
	if err != nil {
		fmt.Println("Call error:", err)
		return
	}

	fmt.Println("Response:", response)
}

// Compare this snippet from src\RPCapi\client.go:
func main() {
	client := InitClient("127.0.0.0:8800")
	var response string
	err := client.HelloServer("client01", &response)
	if err != nil {
		fmt.Println("Call error:", err)
		return
	}
	fmt.Println("Response:", response)
}
