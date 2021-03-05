package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 数组比较
// return
// 1 相等
// 0 长度不相等
// -1 长度相等，值不相等
func arrcmp(a1 []int, a2 []int) int {
	if len(a1) == len(a2) {
		for j, i := range a1 {
			if i != a2[j] {
				return -1
			}
		}
		return 1

	} else {
		return 0
	}
}

func main01() {
	var slice []int = make([]int, 20)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(slice); i++ {
		slice[i] = rand.Intn(40)
	}
	for i := 0; i < len(slice); i++ {
		for j := len(slice) - 1; j > i; j-- {
			if slice[j] > slice[i] {
				slice[j], slice[i] = slice[i], slice[j]
			}
		}
	}
	fmt.Println(slice)

	// 切片
	var s []int = slice[5:10]
	s[3] = 666
	fmt.Println("s == slice[5:10] ?", arrcmp(s, slice[5:10]))

	var a []int = make([]int, 5)
	copy(a, slice[5:10])
	a[4] = 666
	fmt.Println("a == slice[5:10] ?", arrcmp(a, slice[5:10]))
}
