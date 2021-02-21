package main

import "fmt"

func test1(a, b int) int {
	return a + b
}

func add(a int, b int) int {
	return test1(a, b)
}

func main03() {
	a := 20
	b := 10
	fmt.Println(add(a, b))
}
