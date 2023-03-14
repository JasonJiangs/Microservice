package Chatting_Room

import (
	"fmt"
	"net"
)

func main() {
	// init server
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}

	// watch to bubild connection
	connection, err := listen.Accept()
	if err != nil {
		fmt.Println("listen.Accept err:", err)
		return
	}
	fmt.Println("Build connection success!")

	// build connection
	go handler(connection)

}

func handler(connection net.Conn) {
	
}