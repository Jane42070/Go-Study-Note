package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var temp []int = make([]int, 10)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		temp[i] = int(rand.Intn(40))
	}
	fmt.Println(temp)
	// 1. 练习1：从一个整数数组中取出最大的整数，最小整数，平均值
	var min, max, mean, sum = temp[0], temp[0], 0, 0
	for _, i := range temp {
		sum += i
		if min > i {
			min = i
		}
		if max < i {
			max = i
		}
	}
	mean = sum / len(temp)
	fmt.Printf("最大值：%d，最小值：%d\n平均值：%d，和：%d\n", max, min, mean, sum)
	// 3. 反转字符串
	temp_len := len(temp)
	for i := 0; i < temp_len; i++ {
		temp_len--
		temp[i], temp[temp_len] = temp[temp_len], temp[i]
	}
	fmt.Println(temp)
	// 2. 分割所有字符串
	temp_str := []string{"王哥", "黑手", "虎弟", "刀哥"}
	for _, i := range temp_str {
		fmt.Print(i + "|")
	}
}
