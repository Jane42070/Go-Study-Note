package main

import "fmt"

func main05() {
	// 切片名本身就是一个地址
	// 创建的空切片 指向内存地址编号为 0(nil) 的空间
	var slice []int
	const a = 1
	fmt.Printf("%p\n", slice)

	slice = append(slice, 1, 2, 3)
	fmt.Printf("%p\n", slice)

	slice = append(slice, 4, 5, 6)
	fmt.Printf("%p\n", slice)
	// 指针
	var p *int
	fmt.Println(p == nil)
}

func main0501() {
	slice := make([]int, 10)
	fmt.Printf("%p\n", &slice[0])
	fmt.Printf("%p\n", slice)
	fmt.Printf("%p\n", &slice)
	// 在使用 append 进行数据添加时，如果长度超出容量，容量会自动扩容
	// 一般扩容方式为上一次 容量*2 如果超过 1024 字节，每次扩容上一次的 1/4
	// 容量扩容每次都是偶数
	// len(slice) 计算切片的长度
	// cap(slice) 计算切片的容量
	fmt.Printf("size of slice: %d\n", cap(slice))
	slice = append(slice, 1, 2, 3, 4, 5)
	fmt.Printf("size of slice: %d\n", cap(slice))
	slice = append(slice, 1, 2, 3, 4, 5)
	fmt.Printf("size of slice: %d\n", cap(slice))
	slice = append(slice, 1, 2, 3, 4, 5)
	fmt.Printf("size of slice: %d\n", cap(slice))
	slice = append(slice, 1, 2, 3, 4, 5)
	fmt.Printf("size of slice: %d\n", cap(slice))
	slice = append(slice, 1, 2, 3, 4, 5)
	fmt.Printf("size of slice: %d\n", cap(slice))
	slice = append(slice, 1, 2, 3, 4, 5)
	fmt.Printf("size of slice: %d\n", cap(slice))
	slice = append(slice, 1, 2, 3, 4, 5)
	fmt.Printf("size of slice: %d\n", cap(slice))
	slice = append(slice, 1, 2, 3, 4, 5)
	fmt.Printf("size of slice: %d\n", cap(slice))
	slice = append(slice, 1, 2, 3, 4, 5)
	fmt.Printf("size of slice: %d\n", cap(slice))
	slice = append(slice, 1, 2, 3, 4, 5)
	fmt.Printf("size of slice: %d\n", cap(slice))
	slice = append(slice, 1, 2, 3, 4, 5)
	fmt.Printf("size of slice: %d\n", cap(slice))
	slice = append(slice, 1, 2, 3, 4, 5)
	fmt.Printf("size of slice: %d\n", cap(slice))
	slice = append(slice, 1, 2, 3, 4, 5)
	fmt.Printf("size of slice: %d\n", cap(slice))
	slice = append(slice, 1, 2, 3, 4, 5)
	fmt.Printf("size of slice: %d\n", cap(slice))
}
