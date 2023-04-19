package main

import (
	"Entity/Entity"
	"fmt"
	"net"
	"sync"
	"time"
)

var lock sync.Mutex

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
	fmt.Println("Init service...")

	clientAddress := connection.RemoteAddr().String()
	fmt.Println(clientAddress)

	newUser := Entity.User{
		Id:   clientAddress,         // as id
		Name: clientAddress,         // will be renamed when build connection
		Msg:  make(chan string, 10), // need the make space, TODO 管道缓冲和非缓冲区别
	}
	// add user to the global map
	allUsers[newUser.Id] = newUser

	// create a reset timer channel, for telling that user is still online
	var resetTimer = make(chan bool, 1)
	// signal for exit
	var isExit = make(chan bool, 1)
	go watchExit(&newUser, connection, isExit, resetTimer)

	// start a go thread to write message to client
	go writeBackToClient(&newUser, connection)

	// write to global message channel
	lock.Lock()
	loginInfo := fmt.Sprintf("[%s]:[%s] ===> Online", newUser.Id, newUser.Name)
	message <- loginInfo
	lock.Unlock()

	// functional logic for chatting room
	for {
		buf := make([]byte, 1024)
		cnt, err := connection.Read(buf)

		if cnt == 0 {
			fmt.Println("Client exit.")
			lock.Lock()
			isExit <- true
			lock.Unlock()
		}

		if err != nil {
			fmt.Println("conn.Read err:", err)
			return
		}
		fmt.Println("Message from server: ", string(buf[:cnt-1]), " with length ", cnt)

		// ================= 业务逻辑 =================
		// query all users
		uInput := string(buf[:cnt-1])
		if uInput == "who" && len(uInput) == 3 {
			fmt.Println("Query all users.")
			// send all users to client
			var allUsersInfo string
			for _, user := range allUsers {
				allUsersInfo += fmt.Sprintf("[%s]:[%s] ", user.Id, user.Name)
			}
			// send message to user, 如果msg是非缓冲的，会阻塞
			newUser.Msg <- allUsersInfo
		} else if uInput[:6] == "rename" && len(uInput) > 7 { // input: rename|newName
			fmt.Println("Rename user.")
			// rename user
			newName := uInput[7:]
			newUser.Name = newName
			// update user in the global map
			allUsers[newUser.Id] = newUser
			// send message to user, 如果msg是非缓冲的，会阻塞
			newUser.Msg <- "Rename successfully."
		} else {
			// send message to global message channel
			message <- uInput
		}

		resetTimer <- true
		// ================= 业务逻辑 =================

	}
}

func broadcast() {
	fmt.Println("Broadcast go thread init successfully.")

	defer fmt.Println("Broadcast go thread exit.")

	for {
		info := <-message
		fmt.Println("Broadcast go thread get message: ", info)
		// send the message to all users
		for _, user := range allUsers {
			// send message to user, 如果msg是非缓冲的，会阻塞
			user.Msg <- info
		}
	}
}

// write message to client
func writeBackToClient(user *Entity.User, conn net.Conn) {
	fmt.Println("Write go thread init successfully for user: ", user.Name)
	for data := range user.Msg {
		fmt.Println("Write go thread get message: ", data, " for user: ", user.Name)
		// send message to user
		_, err := conn.Write([]byte(data))
		if err != nil {
			fmt.Println("conn.Write err:", err)
			return
		}
	}
}

// go thread for listening exit signal, responsible for closing connection and some other things
func watchExit(user *Entity.User, conn net.Conn, isExit, resetTimer <-chan bool) { // 单向通道
	fmt.Println("Exit go thread init successfully for user: ", user.Name)
	defer fmt.Println("Exit go thread exit successfully for user: ", user.Name)
	for {
		select {
		case <-isExit:
			fmt.Println("Exit go thread with user: ", user.Name)
			delete(allUsers, user.Id)
			message <- fmt.Sprintf("[%s]:[%s] ===> Offline", user.Id, user.Name)
			conn.Close()
			return
		case <-time.After(10 * time.Second):
			fmt.Println("Exit go thread timeout with user: ", user.Name)
			delete(allUsers, user.Id)
			message <- fmt.Sprintf("[%s]:[%s] ===> Offline", user.Id, user.Name)
			conn.Close()
			return

		case <-resetTimer:
			// 重置计数器
			fmt.Println("Reset timer with user: ", user.Name)
		}
	}
}
