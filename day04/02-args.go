package main

import (
	"fmt"
)

func sum(arr ...int) {
	fmt.Print(arr)
}

func main02() {
	sum(1, 2, 3)
}
