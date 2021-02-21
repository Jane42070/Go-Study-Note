package main

import "fmt"

func added(a int) {
	a += 10
	fmt.Println(a)
	if a > 10000 {
		return
	}
	added(a)
}

func main06() {
	added(1)
}
