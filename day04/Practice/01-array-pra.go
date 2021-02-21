package main

import (
	"fmt"
)

func main() {
	var arr [10]int = [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for _, i := range arr {
		fmt.Print(i, " ")
	}
	// 数组赋值
	a := arr
	fmt.Println(a)
	c := 0
	for _, i := range a {
		a[c] = i + arr[c]
		c++
	}
	fmt.Println(a)
}
