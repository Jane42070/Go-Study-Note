package main

import (
	"fmt"
)

func swap(a int, b int) []int {
	a, b = b, a
	return []int{a, b}
}

func main01() {
	a := 10
	b := 20
	temp := swap(a, b)

	fmt.Print(temp)
}
