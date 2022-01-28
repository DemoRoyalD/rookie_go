package main

import "fmt"

func main() {
	// 1. if-else
	var condition = false
	if condition {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
	// 2. witch 与java不同 java: 符合条件 就退出 go: 符合条件就自动break，同时 fallthrough 会强制执行下面的case 不会判断符合不符合case

	switch condition {
	case false:
		fmt.Println("switch : false")
		fallthrough
	case true:
		fmt.Println("switch : true")
		fallthrough
	//case false:
	//	fmt.Println("switch : false")
	//	fallthrough
	default:
		fmt.Println("Default")
	}
}
