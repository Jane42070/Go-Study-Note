package main

import (
	"fmt"
)

func test01(m map[int]string) {
	for k, v := range m {
		fmt.Printf("key = %d, value = %s\n", k, v)
	}
	m[103] = "法师"
	fmt.Print("\n")
}

func main() {
	m := make(map[int]string)
	m[101] = "法师"
	m[102] = "法师"
	// 地址传递 引用传递 形参和实参指向内存中相同的地址段 修改形参会影响实参的值
	test01(m)
	for k, v := range m {
		fmt.Printf("key = %d, value = %s\n", k, v)
	}
}
