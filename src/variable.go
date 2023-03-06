package main

import "fmt"

func main() {
	var name string
	name = "jason"
	var age int
	age = 20
	fmt.Println("name:", name)
	fmt.Printf("name is: %s, age is %d", name, age)

	var gender = "male"
	fmt.Println(gender)

	address := "Beijing"
	fmt.Println(address)

	test(10, "Jason")

	i, j := 1, 2
	fmt.Println(i, j)

	// exchange
	i, j = j, i
	fmt.Println(i, j)

	k := 20
	k++
	fmt.Println(k)

}

func test(a int, b string) {
	fmt.Println(a, b)
}
