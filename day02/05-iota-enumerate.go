// Package main provides ...
package main

import (
	"fmt"
)

func main() {
	// iota 枚举
	// iota 枚举格式如果写在一行中，值相等，如果换行，值加一
	// const (
	//     a    = iota
	//     b    = iota
	//     c, d = iota, iota
	// )
	// const (
	//     a = iota
	//     b
	//     c
	//     d
	// )
	const (
		a = iota
		b = 10
		c = 20
		d
		e
		f = iota
		g
	)
	fmt.Print(a)
	fmt.Print(b)
	fmt.Print(c, d, e, f, g)
}
