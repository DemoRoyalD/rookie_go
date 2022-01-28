package main

import (
	"errors"
	"fmt"
)

func main() {
	// Map
	nameAgeMap := make(map[string]int)
	nameAgeMap["dongchen"] = 20
	nameAgeMap["ly"] = 20
	fmt.Println(nameAgeMap)

	age, ok := nameAgeMap["dongchen"]
	if ok {
		fmt.Println(age)
	}

	// 遍历 Map
	for key, value := range nameAgeMap {
		fmt.Printf(" Key is %v --- value is %v \n", key, value)
	}

	// len
	fmt.Println(len(nameAgeMap))

	result, err := sum(-1, 2)
	fmt.Println(result, err)

}

func sum(num1, num2 int) (int, error) {
	if num1 < 0 || num2 < 0 {
		return 0, errors.New("param must better than zero !")
	}
	return num2 + num1, nil
}
