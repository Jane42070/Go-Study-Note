package main

import (
	"fmt"
)

func main01() {
	var a [10]int
	// 定义数组时，可以为部分元素赋值
	// var arr [10]int = [10]int{1, 2, 3, 4, 5}
	// 定义数组时，可以指定下标赋值
	// var arr [10]int = [10]int{0: 10, 2: 20, 4: 50, 9: 30}
	// 不定参数组赋值
	var arr [10]int = [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := 0; i < 20; i += 2 {
		a[i/2] = i
	}
	fmt.Print(a, "\n")
	// 遍历数组
	// 1
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}
	// 2 range 数组名进行遍历
	// i -> index | replace i to _ as anonymous variable
	// v -> value
	for _, v := range a {
		fmt.Print(v)
	}
	// [1 2 3 4 5 0 0 0 0 0]
	fmt.Println(arr)

	// 赋值必须长度类型一样
	arr = a
	fmt.Print(arr)
}
