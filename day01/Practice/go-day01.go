package main

import "fmt"

// hello Golang
func main() {
	var a int = 16
	b := 20
	a, b = b, a
	fmt.Println("你好，Golang")
	fmt.Printf("there is a shabi variable a:%d\n", a)
	fmt.Print("傻逼 b:", b, "\n")
	fmt.Printf("a = %d, b = %d\n", a, b)
	test01()
}

func test01() {
	var a float32
	fmt.Print("请输入数字：")
	fmt.Scan(&a)
	fmt.Printf("你输入的数字是：%f，类型是：%T, 地址为：%p", a, a, &a)
}
