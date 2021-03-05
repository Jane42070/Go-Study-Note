package main

import (
	"fmt"
)

type Student struct {
	id     int
	name   string
	age    int
	gender byte
}

func Sort(arr []Student) {
	for i := 0; i < len(arr); i++ {
		for j := len(arr) - 1; j > i; j-- {
			if arr[j].age > arr[i].age {
				arr[j], arr[i] = arr[i], arr[j]
			}
		}
	}
	fmt.Println(arr)
}

func main() {
	// 结构体切片
	var arr []Student = []Student{
		Student{1, "唐三藏", 18, 1},
		Student{2, "白龙马", 50, 1},
		Student{3, "猪八戒", 17, 1},
		Student{4, "孙悟空", 19, 1},
		Student{5, "沙悟净", 56, 1},
	}

	fmt.Println(arr)
	arr[1].name = "黑龙马"
	fmt.Println(arr)
	arr = append(arr, Student{6, "玉帝老儿", 9000, 1})
	// 结构体数组作为函数参数是值传递
	// 结构体切片作为函数参数是地址传递
	Sort(arr)
	fmt.Println(arr)
}
