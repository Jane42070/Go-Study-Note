// Package main provides ...
package main

import (
	"fmt"
	"math"
)

func main02() {
	// 变量的定义和使用
	// var 变量名 数据类型 = 值
	var sun int = 50

	// 变量在程序运行过程中 值可以发生改变
	// 表达式
	sun += 25

	fmt.Print(sun)
}

func main0201() {
	// 如果没有赋值 默认值为 0
	// var sun int
	var value float64 = 2

	// sun = value * value * value * value * value * value * value * value
	// 可以使用系统提供的包
	var sum float64 = math.Pow(value, 10)

	fmt.Print("sum = ", sum)
}
