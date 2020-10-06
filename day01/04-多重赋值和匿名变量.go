// Package main provides ...
package main

import "fmt"

func main04() {
	// a := 10
	// b := 20
	// c := 30

	// 多重赋值
	// 变量个数和值的个数要一一对应
	a, b, c := 10, 20, "Go"

	fmt.Println(a, b, c)

	// 如果在多重赋值时有新定义的变量，可以使用自动推导类型
	// a, b, c, d := 20, 30, "GoLang", "Hello"

	// _表示匿名变量，不接收数据
	_, c, d := 110, "Hello", "World"

	// fmt.Println(_)
	fmt.Println(a, b, c, d)
}
