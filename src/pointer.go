package main

import "fmt"

func main() {
	// have garbage collector

	name := "Jason"
	ptr := &name
	fmt.Println(*ptr)
	fmt.Println(ptr)

	ptr2 := new(string)
	*ptr2 = "Duke"
	fmt.Println(*ptr2)
	fmt.Println(ptr2)

	city := "Baltimore"
	fmt.Println(city)
	fmt.Println(&city)
	fmt.Println(testPtr(city))
	fmt.Println(*testPtr(city))

	if ptr == nil {
		fmt.Println("name is null")
	} else {
		fmt.Println("Not null")
	}

}

func testPtr(city string) *string {
	return &city
}
