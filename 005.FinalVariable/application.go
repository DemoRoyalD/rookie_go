package main

import "fmt"

func main() {

	// 1. 常量  不能modify 的变量
	const MODE string = "Can't Modify"
	//MODE = "b"
	fmt.Println(MODE)

	// 2. iota 特殊常量  iota 在 const关键字出现时将被重置为 0(const 内部的第一行之前)，const 中每新增一行常量声明将使 iota 计数一次
	//const (
	//	FIRST = iota
	//	SECOND = iota
	//	THIRD = iota
	//)

	//const (
	//	FIRST = iota
	//	SECOND
	//	THIRD
	//)

	const (
		FIRST = iota
		SECOND
		THIRD
		FOURTH = "ha"
		FIVE   = 100
		SIX    // = 100
		SEVEN  = iota
	)
	fmt.Println(FIRST, SECOND, THIRD, FOURTH, FIVE, SIX, SEVEN)
}
