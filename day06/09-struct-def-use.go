package main

import (
	"fmt"
)

// type 结构体名 struct {
// 成员
// }
type student struct {
	name   string
	age    int
	height float64
	gender byte
}

func main() {
	// 定义结构体变量，符合类型
	// var 变量名 结构体名
	var stu student
	var stu1 student = student{"小明", 'f', 1.68, 20}
	stu2 := student{"孙殿明", 'm', 1.99, 35}
	stu3 := student{gender: 'f', name: "李华"}

	// 结构体成员赋值
	stu.name = "大名"
	stu.gender = 'f'
	stu.height = 1.77
	stu.age = 20
	fmt.Println(stu)
	fmt.Println(stu1)
	fmt.Println(stu2)
	fmt.Println(stu3)

	// 结构体名本身指向第一个成员的地址
	fmt.Printf("%p\n", &stu)
	fmt.Printf("%p\n", &stu.name)
	// string 类型需要和结构体最大的数据类型进行对齐
	fmt.Printf("%p\n", &stu.age)
	fmt.Printf("%p\n", &stu1)
	fmt.Printf("%p\n", &stu2)
	fmt.Printf("%p\n", &stu3)

	if stu1 != stu2 {
		fmt.Println("不相等")
	}

	if stu.height > stu1.height {
		fmt.Println("stu > stu1")
	}
}
