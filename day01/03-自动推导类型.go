// Package main provides ...
package main

import (
	"fmt"
)

func main03() {
	// var a = 20
	// var c int = 40

	a := 20.0
	b := 40
	c := "Deutsch"
	d := 12.345678

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(a + d)
	fmt.Println(a + 3.14159265)
	fmt.Println("--------------")

	fmt.Println("Swap each variables' values")
	swap(a, d)
}

// 交换两个变量的值
func swap(a, b float64) {
	temp := a
	a = b
	b = temp

	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
}
