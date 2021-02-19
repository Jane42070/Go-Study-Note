// Package main provides ...
package main

import (
	"fmt"
)

func main() {
	a := 10
	// 取地址
	fmt.Println(&a)
	// 取值
	fmt.Println(*(&a))
}
