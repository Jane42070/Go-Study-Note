// Package main provides ...
package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 其他类型转换成字符串添加到字符切片里面
	slice := make([]byte, 0, 1024)
	fmt.Println(slice)
	slice = strconv.AppendBool(slice, true)
	slice = strconv.AppendInt(slice, 123, 2)
	slice = strconv.AppendFloat(slice, 3.1415926, 'f', 4, 64)
	slice = strconv.AppendQuote(slice, "hello")
	fmt.Println(string(slice))
}
