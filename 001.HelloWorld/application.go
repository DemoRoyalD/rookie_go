package main

import "fmt"

func main() {
	fmt.Println("Hello World! ")
	var name string
	// 阻塞线程
	fmt.Scanln(&name)
	fmt.Println(name)

	//	Command  : go run application.go
	// Generate Binary File: go build   /
}
