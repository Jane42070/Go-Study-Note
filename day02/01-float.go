// Package main provides ...
package main

import (
	"fmt"
)

func main() {
	// 浮点型数据 分为单精度浮点型 float32 （小数位数为 7 位） 双精度浮点型 float64 （小数位数为 15 位）
	// float64 比 float32 更精准
	var a float64
	fmt.Print("Input a float number:")
	fmt.Scan(&a)
	// 保留六位小数数据会更精准
	fmt.Printf("Variable a = %f, typeof(a) = %T\n", a, a)
	// 通过自动推导类型创建的浮点型变量 默认为 float64
	b := 123.456
	fmt.Printf("%T\n", b)
}
