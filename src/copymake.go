package main

import "fmt"

func main() {
	cities := [5]string{"Baltimore", "Boston", "NYC", "Columbia", "Beijing"}

	city1 := cities[0:3]
	fmt.Println(city1)

	city1[2] = "changed"
	fmt.Println(cities)
	fmt.Println(city1)

	fmt.Println(cities[:4])
	fmt.Println(cities[4:5])
	fmt.Println(cities[:])

	sub1 := "helloworld"[5:7]
	fmt.Println(sub1)

	str2 := make([]string, 10, 20)
	fmt.Println(str2)
	fmt.Println(len(str2))
	fmt.Println(cap(str2))

	cityCopy := make([]string, len(city1))
	copy(cityCopy, city1[:])
	fmt.Println(cityCopy)
}
