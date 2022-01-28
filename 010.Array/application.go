package main

import "fmt"

func main() {

	// 1. introduce array
	var intArray [10]float32
	// 2. initial array
	var balance = [5]float64{1, 3, 4, 5, 2}
	balance2 := [5]float64{1, 2, 3, 4, 5}
	// 不固定 的话 [...]
	fmt.Println(intArray)
	fmt.Println(balance)
	fmt.Println(balance2)

	// range
	for _, value := range balance {
		fmt.Println(value)
	}
	for index, value := range balance {
		fmt.Println(index, value)
	}

}
