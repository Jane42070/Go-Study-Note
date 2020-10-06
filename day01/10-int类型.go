// Package main provides ...
package main

import (
	"fmt"
	"math"
)

func main1001() {
	// int 类型会根据操作系统位数不同在内存中占的字节也不同
	// 64 位系统为 int64 8 字节
	// 32 位系统为 int32 4 字节
	var a int = 10

	// bit 是计算机最小和储存单元，一个位有两个状态，通电（开）、断电（关）状态
	// 用 0 或 1 表示
	// 对应的是二进制数据
	// Byte 字节 一个字节等于八个位 1B = 8bit
	// 1KB = 1024B = 8192bit
	// 1MB = 1024KB
	// 1GB = 1024MB
	// 1TB = 1024GB
	// 1PB = 1024TB

	// 8B = 64bit
	// 1 符号位 63 数字位
	// - 2^63 ~ 2^63 - 1
	// 0 占了其中一位，所以最大值要 -1

	fmt.Println("int64取值范围：", math.Pow(-2, 63), "~", math.Pow(2, 63)-1)
	fmt.Println("int32取值范围：", math.Pow(-2, 31), "~", math.Pow(2, 31)-1)

	fmt.Println(a)
}

func main1002() {
	var a uint = 10
	// 如果出现了负数，会溢出，从最小值变成最大值
	a = a - 20

	fmt.Println("uint取值范围：", 0, "~", math.Pow(2, 64)-1)
	fmt.Println(a)
}

func main() {
	var a int = 10
	var b int64 = 60

	// 定义 int 和 int64 是需要使用类型转换才可以计算
	fmt.Println(a + b)
}
