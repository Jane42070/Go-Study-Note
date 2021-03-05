package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := "hello world"

	// 字符串转成字符切片
	slice := []byte(str)
	fmt.Printf("%s, len = %d\n", slice, len(slice))
	// 切片转为字符串
	fmt.Println(string(slice))
	// 将其他类型转化成字符串
	str = strconv.FormatBool(true)
	fmt.Println(str)
	// 在计算机中进制可以表示 2-36 进制
	str = strconv.FormatInt(123, 24)
	fmt.Println(str)
	str = strconv.FormatFloat(3.1415926, 'f', 6, 64)
	fmt.Println(str)
	// iota itoa
	// Itoa is equivalent to FormatInt(int64(i), 10).
	str = strconv.Itoa(123)
	fmt.Println(str)

	// string 转成布尔
	b, err := strconv.ParseBool("真")
	if err != nil {
		fmt.Println("类型转换出错")
	} else {
		fmt.Println(b)
	}

	// 当成 2 进制转换成 10 进制为字符串
	value, _ := strconv.ParseInt("101", 2, 64)
	fmt.Println(value)
}
