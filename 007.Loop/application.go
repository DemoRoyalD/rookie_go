package main

import "fmt"

func main() {

	// 1. for continue break

	var index = 10
	for mark := 0; mark < index; mark++ {
		if mark == 5 {
			continue
		}
		if mark == 8 {
			break
		}
		fmt.Println("Mark  = ", mark)
	}

}
