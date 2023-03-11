package main

import "fmt"

func main() {
	var idNames map[int]string // 不能直接赋值

	idNames = make(map[int]string, 10)

	idNames[0] = "Duke"
	idNames[1] = "JHU"

	for key, value := range idNames {
		fmt.Println(key, value)
	}

	// no out of bound
	fmt.Println(idNames[9])
	fmt.Println(idNames[20])

	idNumbers := make(map[int]float64, 10)
	fmt.Println(idNumbers[11])

	value, exist := idNames[1]
	if exist {
		fmt.Printf("id=%d exist: %s\n", 1, value)
	}

	delete(idNames, 0)
	delete(idNames, 100)
	fmt.Println(idNames)
}
