package main

import (
	"Entity/Entity"
	"fmt"
	"net"
)

// global map to save all users
var allUsers = make(map[string]Entity.User)

// global channel for message
var message = make(chan string, 10)

func main() {
	// init server
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}
	fmt.Println("Server successfully init!")

	// 启动全局唯一的go程，负责监听message通道，发送给所有用户
	go broadcast()

	for {
		fmt.Println("主go程监听中============")

		// watch to build connection
		connection, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept err:", err)
			return
		}

		// build connection
		go handler(connection)

	}

}

// do service function
func handler(connection net.Conn) {
	for {
		fmt.Println("Init service...")

		clientAddress := connection.RemoteAddr().String()
		fmt.Println(clientAddress)

		newUser := Entity.User{
			Id:   clientAddress,     // as id
			Name: clientAddress,     // will be renamed when build connection
			Msg:  make(chan string), // need the make space
		}
		// add user to the global map
		allUsers[newUser.Id] = newUser
		// write to global message channel
		loginInfo := fmt.Sprintf("[%s]:[%s] ===> Online", newUser.Id, newUser.Name)
		message <- loginInfo

		buf := make([]byte, 1024)
		cnt, err := connection.Read(buf)
		if err != nil {
			fmt.Println("conn.Read err:", err)
			return
		}
		fmt.Println("Data from server: ", string(buf[:cnt]), " with length ", cnt)

	}
}

func broadcast() {
	fmt.Println("Broadcast go thread init successfully.")
	info := <-message
	// send the message to all users
	for _, user := range allUsers {
		user.Msg <- info
	}
}
