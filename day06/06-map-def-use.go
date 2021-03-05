// Package main provides ...
package main

import (
	"fmt"
)

func main0601() {
	// map[key]value
	// map『键类型』值类型
	// map 存储的方式不是顺序存储的
	var m map[int]string = map[int]string{101: "法师", 251: "张超", 666: "怡红"}
	// fmt.Println(m)
	for k, v := range m {
		fmt.Printf("key = %d, value = %s\n", k, v)
	}
}

func main() {
	m := make(map[string]int)
	// map 长度是自动扩容的
	// map 中 key 唯一，值重复
	fmt.Printf("len = %d\n", len(m))
	m["法师"] = 101
	fmt.Printf("len = %d\n", len(m))
	for k, v := range m {
		fmt.Printf("key = %s, value = %d\n", k, v)
	}
	fmt.Printf("%p\n", m)
}
