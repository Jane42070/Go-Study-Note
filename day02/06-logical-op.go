// Package main provides ...
package main

import (
	"fmt"
)

func main() {
	a := 1
	b := 0
	// 逻辑与高于逻辑或
	fmt.Print(a > b && b > a || a < 0)
}
