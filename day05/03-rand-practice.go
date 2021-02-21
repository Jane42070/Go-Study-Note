package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main03() {
	arr := [20]int{0}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(20)
	}
	fmt.Println(arr)
}
