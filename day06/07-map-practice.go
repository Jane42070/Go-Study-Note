package main

import (
	"fmt"
)

func main() {
	// 通过键盘输入一个英文字符串 统计每个字母出现的次数
	// helloworld
	var m map[string]int = make(map[string]int)
	fmt.Println(m)
	str := "helloworld"
	fmt.Println(str)
	for i := 0; i < len(str); i++ {
		m[string(str[i])]++
	}
	for k, v := range m {
		fmt.Printf("key = %s, value = %d\n", k, v)
	}

	var m1 map[string][3]int = make(map[string][3]int)
	m1["小宁"] = [3]int{40, 50, 70}
	for k, v := range m1 {
		fmt.Printf("key = %s, value = %d\n", k, v)
		for _, i := range v {
			fmt.Println(i)
		}
	}

	value, ok := m["小宁"]
	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("Not found")
	}

	delete(m1, "小宁")
	fmt.Println(m1)
}
