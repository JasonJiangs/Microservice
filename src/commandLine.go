package main

import (
	"fmt"
	"os"
)

func main() {

	// go build commandLine.go
	// ./commandLine.exe hello world
	cmds := os.Args
	//for key, cmd := range cmds {
	//	fmt.Println(key, cmd)
	//}
	//fmt.Println(len(cmds))

	if len(cmds) < 2 {
		fmt.Println("Input again")
		return
	}

	switch cmds[1] {
	case "hello":
		fmt.Println("hello")
		fallthrough
	case "world":
		fmt.Println("world")
	default:
		fmt.Println("default value.")
	}

}
