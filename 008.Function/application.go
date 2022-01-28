package main

import (
	"fmt"
	"math"
)

func main() {

	// 1.函数
	param1 := 1
	param2 := 2
	var max = max(param1, param2)
	fmt.Println("Max = ", max)
	str1 := "First"
	str2 := "Second"
	fmt.Printf("First = %v ; Secod = %v \n", str1, str2)
	str1, str2 = swap(str1, str2)
	fmt.Printf("First = %v ; Secod = %v \n", str1, str2)

	fmt.Printf("First = %v ; Secod = %v \n", str1, str2)
	quoteSwap(&str1, &str2)
	fmt.Printf("First = %v ; Secod = %v \n", str1, str2)

	// 2. function as param
	getSquareRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}
	fmt.Println(getSquareRoot(9))

	// 3. close package
	nextNumber := getSequence()
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	nextNumber1 := getSequence()
	fmt.Println(nextNumber1())
	fmt.Println(nextNumber1())
	fmt.Println(nextNumber1())

	var circle Circle
	circle.radius = 10
	fmt.Println("The area of circle is ", circle.getArea())
}

// 1.Simple function
func max(num1, num2 int) int {
	var result int
	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}

// 2. Many return
func swap(str1, str2 string) (string, string) {
	return str2, str1
}

// 3. quote value transmit
func quoteSwap(x *string, y *string) {
	var temp string
	temp = *x /* 保持 x 地址上的值 */
	*x = *y   /* 将 y 值赋给 x */
	*y = temp /* 将 temp 值赋给 y */
}

// 4. close package
func getSequence() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

// 5. method
type Circle struct {
	radius float64
}

func (circle Circle) getArea() float64 {
	return 3.14 * circle.radius * circle.radius
}
