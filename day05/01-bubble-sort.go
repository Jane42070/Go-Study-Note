package main

import (
	"fmt"
	"math/rand"
)

func bubble_sort(ascending string, arr [100]int) [100]int {
	time := 0
	for i := 0; i < len(arr)-1; i++ {
		// 这是什么排序
		for j := len(arr) - 1; j > i; j-- {
			if arr[j] > arr[i] {
				arr[j], arr[i] = arr[i], arr[j]
				time++
			}
		}
		// 冒泡排序
		// for j := 0; j < len(arr)-i-1; j++ {
		//     if arr[j] < arr[j+1] {
		//         arr[j], arr[j+1] = arr[j+1], arr[j]
		//         time++
		//     }
		// }
	}
	fmt.Printf("比较了:%d次\n", time)
	return arr
}

func main01() {
	var arr [100]int = [100]int{0}
	rand.Seed(196)
	for i := 0; i < 100; i++ {
		arr[i] = rand.Intn(100)
	}
	fmt.Println(arr)
	fmt.Println(bubble_sort("True", arr))
}
