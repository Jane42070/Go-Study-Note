package main

import "fmt"

func main04() {
	// 数组定义
	// var 数据名 [元素个数]数据类型

	// 切片定义
	var slice []int = []int{1, 2, 3, 4, 5, 6}
	// 切片
	fmt.Println(slice[1:5])
	// 添加数据
	slice = append(slice, 7, 8, 9)
	fmt.Println(slice)

	var slice01 []int
	// 不能此方法添加数据
	// slice01[0] = 123
	// 需要通过 append
	slice01 = append(slice01, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(len(slice01), slice01)

	// 在定义切片时可以指定长度 make([]string, 0)
	var slice02 []int = make([]int, 10)
	fmt.Println(slice02)
	fmt.Println(len(slice02))

	slice02 = append(slice02, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(slice02)
}
