package main

import (
	"fmt"
	"strings"
)

func main() {
	// 查找
	// func strings.Contains(s string, substr string) bool
	// func strings.Index(s string, substr string) int

	// 分割组合
	// func strings.Join(elems []string, sep string) string
	// func strings.Split(s string, sep string) []string

	// 重复和替换
	// func strings.Replace(s string, old string, new string, n int) string
	// func strings.Repeat(s string, count int) string

	// 去掉内容
	// func strings.Trim(s string, cutset string) string
	// func strings.Trim(s string, cutset string) string

	str1 := "hello world"
	str2 := "llo"
	ret := strings.Contains(str1, str2)
	if ret {
		fmt.Println("找到了")
	} else {
		fmt.Println("没找到")
	}
	// 字符串切片
	slice := []string{"132", "7293", "4079"}
	// 字符串连接
	// Join
	str3 := strings.Join(slice, "-")
	fmt.Println(str3)
	// 字符串索引，返回下标
	// Index
	i := strings.Index(str1, str2)
	fmt.Println(i)
	str := "性感法师在线讲课"
	// 重复字符串
	// Repeat
	str = strings.Repeat(str, 3)
	fmt.Println(str)
	// 字符串替换
	// Replace, 目标字符串，替换字符串，替换的字符，多少次
	// 用作于屏蔽敏感词汇
	str = strings.Replace(str, "性感", "**", -1)
	fmt.Println(str)

	// 字符串拼接
	// Split 分割
	// Join 拼接
	str = "1300-199-1433"
	slice = strings.Split(str, "-")
	slice1 := strings.Join(slice, "")
	fmt.Println(slice1)

	// 去掉字符串头尾的字符串
	// Trim
	str = "      A u Ok       "
	str = strings.Trim(str, " ")
	fmt.Println(str)

	// 去掉空格，把剩余的元素放入切片中
	str = "      A u Ok       "
	slice2 := strings.Fields(str)
	fmt.Println(slice2)
	for _, i := range slice2 {
		fmt.Printf("%s\n", i)
	}
}
