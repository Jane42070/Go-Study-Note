// Package main provides ...
package main

import (
	"fmt"
)

func test(s []int) {
	fmt.Printf("add of slice : %p\n", s)
	// 如果在函数中使用 append 进行数据添加时，形参的地址改变不会指向实参的地址
	// go 会开辟新的空间存储追加后的数组
	// 如果没有指向的空间则释放
	s = append(s, 10, 11, 12, 13)
	fmt.Printf("add of slice : %p\n", s)
	fmt.Println(s)
}

func main() {
	// 切片名本身就是一个地址
	// 存储在堆区 heap
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("add of slice : %p\n", slice)
	// 地址传递
	test(slice)
	fmt.Printf("add of slice : %p\n", slice)
	fmt.Println(slice)
}
