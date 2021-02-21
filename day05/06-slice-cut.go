package main

import "fmt"

func main06() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s := slice[3:8]
	// 切片名[起始位置:结束位置+1:容量] 切片名[low:high;max]
	fmt.Print(s, "\n")

	fmt.Println(len(s))
	fmt.Println(cap(s))

	// 容量要大于等于长度
	b := slice[3:6:10]
	fmt.Println(b)

	fmt.Println(len(s))
	fmt.Println(cap(s))
	b = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	fmt.Println(b)
}

func main0601() {
	// 截取后的切片还是原始切片中的一块内容，浅拷贝，影响原始切片的值
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s := slice[3:8]
	fmt.Println(s)
	s[0] = 555
	fmt.Println(s)
	fmt.Println(slice)
}
