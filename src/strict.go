package main

import "fmt"

type stu struct {
	name   string
	age    int
	gender string
	score  float64
}

type MyInt int

func main() {
	var i, j MyInt

	i, j = 10, 20
	fmt.Println(i + j)

	jim := stu{
		name:   "Jim",
		age:    20,
		gender: "Male",
		score:  100,
	}
	fmt.Println(jim)

	jimPtr := &jim
	fmt.Println(jimPtr)
	fmt.Println(*jimPtr)
	fmt.Println(jimPtr.name)
	fmt.Println(&jimPtr.name)
}
