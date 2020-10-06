// Package main provides ...
// 行注释
/*
块注释
可以注释多行
注释内容不参与程序编译
*/
// 文件所属的包 在 go 语言中 主函数所在的包一定是 main
package main

import (
	// 标准输入输出的包 / 库
	"fmt"
)

// func 函数格式 main 函数名 主函数 程序有且只有一个主函数
func main01() {
	// fmt 包 Println 打印并换行
	fmt.Println("Hello, world")
	fmt.Println("蹇棋林")
}
