// Package main provides ...
package main

import (
	"fmt"
)

func main() {
	// byte 也是 uint8
	var a byte = 'a'
	var b uint8 = 'b'
	type shabi byte
	var c shabi = '0'
	fmt.Println(a)
	fmt.Printf("%c\n", a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	fmt.Printf("%c\n", b)
	fmt.Printf("%T\n", b)
	// 字符 0 对应的 ASCII 值为48
	fmt.Printf("variable c is %c, %c ASCII code is %d\n", c, c, c)
	fmt.Printf("%c\n", 48)
}
