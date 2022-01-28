package main

import "fmt"

func main() {
	// 1.切片  类似于数组，不同于数组，数组需要固定长度，切片不需要
	// var identifier []type
	// var slice1 []type = make([]type, len)
	// slice1: = make([]type, len)
	s := []int{1, 2, 3}
	fmt.Println(s)
}
