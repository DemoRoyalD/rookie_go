package main

import "fmt"

func main() {
	// 1.切片  类似于数组，不同于数组，数组需要固定长度，切片不需要
	// var identifier []type
	// var slice1 []type = make([]type, len)
	// slice1: = make([]type, len)
	s := []int{1, 2, 3}
	fmt.Println(s)

	// 切片引用的还是数组，所以修改后，引用对象也会被修改
	array := [5]string{"a", "b", "c", "d", "e"}
	slice := array[2:5]
	fmt.Println(slice)
	slice[1] = "f"
	fmt.Println(array)

	// append
	slice2 := append(slice, "f")
	fmt.Println(slice2)

}
