package main

import "fmt"

func main() {
	usage := `./a.out <option>
			-h help			
			-a xxxx`
	fmt.Println(usage)

	fmt.Println(len(usage))

	for i := 0; i < len(usage); i++ {
		fmt.Printf("i: %d, v: %c\n", i, usage[i])
	}

	i, j := "hello", "World"
	fmt.Println("i+j=", i+j)

	// unchangeable for const
	const address = "Baltimore"
	//address = "Maryland"
	fmt.Println(address)

}
