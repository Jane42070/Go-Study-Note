package main

import (
	"fmt"
)

func main() {
	var score int
	fmt.Scan(&score)
	if score > 700 {
		fmt.Print("擅长鼻")
	} else if score > 100 {
		fmt.Print("有点擅长鼻")
	} else {
		fmt.Print("不擅长鼻")
	}
}
