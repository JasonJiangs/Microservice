package main

import "fmt"

// 内存逃逸

/*
go build -help
go help build
--gcflag
go build -o qqq.exe --gcflags "-m -m -l" qqq.go > qqq.txt 2>&1
grep -E "name|city" qqq.txt --color
*/

func main() {
	p1 := testPtr1()
	fmt.Println(*p1)
}

func testPtr1() *string {
	name := "JHU"
	p := &name
	fmt.Println(*p)

	city := "baltimore"
	ptr := &city
	return ptr
}
