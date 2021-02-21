package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main02() {
	// 导入头文件
	// 创建随机数种子
	// 创建随机数

	// 创建随机数种子
	rand.Seed(time.Now().UnixNano())
	fmt.Println(time.Now().UnixNano())
	fmt.Println(rand.Int63n(200))
}
