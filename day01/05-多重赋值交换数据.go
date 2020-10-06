// Package main provides ...
package main

import "fmt"

func main05() {
	a, b := 10, 20

	// 可以通过多重赋值进行交换
	a, b = b, a

	fmt.Println(a, b)
}
