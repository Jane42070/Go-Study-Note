// Package main provides ...
package main

import "fmt"

func main0701() {
	// 变量声明
	var a, b, c int

	// 通过键盘为变量赋值
	// & 是一个运算符 取地址运算符
	// fmt.Scanf = fmt.Scan + formatter
	// fmt.Scanf("%d %d %d", &a, &b, &c)
	// 当输入非变量类型的常量时，不会报错，默认初值为 0
	fmt.Scan(&a, &b, &c)
	// 内存地址
	// 0xc0000b4008 0xc0000b4010 0xc0000b4018

	fmt.Printf("%d %d %d\n", a, b, c)
	fmt.Printf("%p %p %p", &a, &b, &c)
}

func main0702() {
	// 如果多个变量类型相同可以一次声明多个变量
	// var a, b int
	var a int
	var b int

	// 空格或回车 表示一个输入接收的结束
	fmt.Scan(&a, &b)
	fmt.Printf("%d %d\n", a, b)
}

func main0703() {
	var a int
	var b string

	// 在接收字符串时使用空格作为分隔
	fmt.Scanf("%d%s", &a, &b)

	fmt.Println(a)
	fmt.Printf("===%s===", b)
}

func main0704() {
	var c, m, e float64

	fmt.Scan(&c, &m, &e)

	sum := c + m + e

	fmt.Println("总成绩：", sum)
	// 两个整型数据相除，得到的结果也是整型，小数部分的数据舍去
	fmt.Println("平均成绩：", sum/3)
	fmt.Printf("%p %p %p", &c, &m, &e)
}
