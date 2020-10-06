// Package main provides ...
package main

import "fmt"

func main06() {
	// fmt.Println("Hello, world!")
	// fmt.Print("性感荷官在线发牌")
	// fmt.Print("澳门赌场上线啦")
	a := 10
	b := 1.2345678

	fmt.Println(a)
	// format output
	// 格式化输出
	// %d 是一个占位符 表示整型数据
	// %f 是一个占位符 表示输出一个浮点型数据 %f 默认保留六位小数
	// 因为浮点型数据不是精准的数据，6 位是有效地
	// %.2f 保留两位小数，会对后一位小数进行四舍五入
	fmt.Printf("%d %.3f\n", a, b)

	c := "你瞅啥"
	// %s 是一个占位符 表示输出一个字符串类型
	fmt.Printf("%s\n", c)
}
