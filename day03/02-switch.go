package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Scan(&a)
	switch a {
	case 1:
		fmt.Print("111111")
	case 2:
		fmt.Print("222222")
	default:
		fmt.Print(a, a, a, a, a, a, "\n")
	}
	switch a > 200 {
	case true:
		fmt.Print("111111\n")
		// 继续执行下面的代码
		fallthrough
	case false:
		fmt.Print("222222")
	}
}
