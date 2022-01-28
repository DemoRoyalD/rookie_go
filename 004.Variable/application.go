package main

import "fmt"

// 全局变量
var (
	a int
	b bool
)

func main() {

	// 1. 声明变量   var identifier type
	var v1, v2 string
	v1 = "test1"
	v2 = "test2"
	fmt.Println(v1, v2)

	v3 := 1 // 等同于 var v3 int ; v3 = 1
	fmt.Println(v3)
	fmt.Println(a)
	fmt.Println(b)

	// 2.值类型和引用类型
	// int float bool string 基本类型都属于值类型 j = i 其实是在内存中将i的值进行了拷贝  &i 取得是物理地址 *i 引用对象,获取的是值
	var quote = 1
	var copy = &quote
	fmt.Println("copy value = ", copy)
	fmt.Println("copy physical address = ", &copy)
	fmt.Println("quote physical address =", &quote)
	*copy = 2
	fmt.Println("copy value = ", *copy)
	fmt.Println("quote value =", quote)

}
