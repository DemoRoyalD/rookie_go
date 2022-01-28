package main

import "fmt"

func main() {

	// 1.格式化字符串
	// %d 表示整数  %s 表示字符串
	var stockCode = 123
	var endDate = "2022-01-28"
	var url = "Code = %d endDate = %s"
	var targetUrl = fmt.Sprintf(url, stockCode, endDate)
	fmt.Println(targetUrl)

}
