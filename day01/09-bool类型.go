// Package main provides ...
package main

import "fmt"

func main() {
	// 布尔类型
	// 一般用于条件判断
	// 默认值为 false
	var a bool
	var b bool = true

	// 通过自动推导类型创建布尔类型变量
	c := false
	a = true

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	// fmt.Println(a)
	// %T 是一个占位符 表示输出一个变量对应的数据类型
	fmt.Printf("%T\n", c)

	d := 10
	fmt.Printf("%T\n", d)

	e := 3.14159265
	fmt.Printf("%T\n", e)

	f := "花姑凉"
	fmt.Printf("%T\n", f)
}
