package main

import "fmt"

func test1(a int, b int, c string) (int, string, bool) {
	return a + b, c, true
}

func test2(a int, b int, c string) (i int, s string, bo bool) {
	i = a + b
	s = c
	bo = true
	return
}

func main() {
	v1, v2, v3 := test1(10, 20, "hey")
	fmt.Println(v1, v2, v3)

	v4, v5, v6 := test2(10, 20, "hey")
	fmt.Println(v4, v5, v6)
}
