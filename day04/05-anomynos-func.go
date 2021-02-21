package main

import (
	"fmt"
)

func main05() {
	a := 10
	b := 20

	// 在函数内部定义一个匿名函数
	// f 为函数类型
	var f func(a int, b int) int = func(a int, b int) int {
		return a + b
	}
	a = 100
	b = 200
	fmt.Println(f(a, b))
	fmt.Printf("%T\n", f)
	fmt.Println("Address of f:", &f)
}
