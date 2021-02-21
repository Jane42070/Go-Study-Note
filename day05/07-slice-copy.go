package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(slice)
	s := make([]int, 10)
	fmt.Println(s)
	// 使用 copy 进行拷贝，在内存中存储两个独立的切片内容，互不影响
	copy(s, slice)
	fmt.Println(s)

	s[0] = 222
	fmt.Println(s)
	fmt.Println(slice)

	fmt.Printf("&s = %p\n&slice = %p\n", &s, &slice)
}
