// Package main provides ...
package main

import (
	"fmt"
)

func main() {
	var a string = "hello world"
	b := "hello str"
	fmt.Printf("a is %s, b is %s\n", a, b)
	c := "蹇棋林"
	// 在 go 语言中，一个汉字占 3 个字符
	fmt.Printf("c len = %d\n", len(c))

	str1 := "澳门在线赌场上线了"
	str2 := "性感荷官在线发牌"

	str3 := str1 + str2
	fmt.Print(str3)
}
