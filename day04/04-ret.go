package main

import "fmt"

func arr_print(sum int, arr ...int) int {
	return sum
}

func arr_sum(arr ...int) int {
	var sum int = 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	return arr_print(sum)
}

func main04() {
	sum := arr_sum(1, 2, 3, 4, 5)
	fmt.Print(sum)
}
