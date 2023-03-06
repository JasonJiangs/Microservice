package main

import "fmt"

func main() {
	//var nums = [10]int{1, 2, 2, 4}
	//var nums [10]int = [10]int{1, 2, 2, 4}
	nums := [10]int{1, 2, 2, 4}
	fmt.Println(nums)

	for i := 0; i < len(nums); i++ {
		fmt.Println(i, nums[i])
	}

	fmt.Println("______")

	for i, num := range nums {
		fmt.Println(i, num)
	}

	fmt.Println("______")
	for _, num := range nums {
		fmt.Println(num)
	}

	fmt.Println("______")
	for _, _ = range nums {
		fmt.Println("dd")
	}

	// make
	cities := []string{"N", "M", "J", "Z"}
	cities = append(cities, "K")
	cities = append(cities, "K")
	fmt.Println(cities)
	fmt.Println(len(cities))
	fmt.Println(cap(cities))
}
