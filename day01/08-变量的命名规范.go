// Package main provides ...
package main

import "fmt"

func main() {
	// 1. 允许使用字母 数字 下划线
	// 2. 不允许使用关键字
	// 3. 不允许使用数字开头和单独的下划线_
	// 4. 区分大小写
	// 5. 见名知意

	var ___ int = 10
	var _abc int = 30
	var ABC int = 40
	var abc int = 50

	// 大驼峰命名
	// RoleLevel
	// 小驼峰命名
	// roleLevel
	// Linux 命名
	// role_level

	fmt.Println(___)
	fmt.Println(_abc)
	fmt.Println(ABC)
	fmt.Println(abc)
}
