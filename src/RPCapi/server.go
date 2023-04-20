package main

import (
	"fmt"
	"net"
	"net/rpc"
)

// define a struct
type Server struct {
}

// define a method for the struct
func (s *Server) HelloServer(name string, response *string) error {
	*response = "hello:" + name
	return nil
}

func main() {
	// register rpc server, the first parameter is the name of the service, the second parameter is the service instance
	err := rpc.RegisterName("server01", new(Server))
	if err != nil {
		fmt.Println("RegisterName error:", err)
		return
	}

	// set listen address
	listen, err := net.Listen("tcp", "127.0.0.1:8800")
	if err != nil {
		fmt.Println("Listen error:", err)
		return
	}
	defer listen.Close()

	fmt.Println("Start listen on", listen.Addr().String())
	// build a connection
	conn, err := listen.Accept()
	if err != nil {
		fmt.Println("Accept error:", err)
		return
	}
	defer conn.Close()
	fmt.Println("Accept a connection from", conn.RemoteAddr().String())

	// bind the connection to rpc
	rpc.ServeConn(conn)

}
