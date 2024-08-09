package main

import (
	"fmt"
	"strconv"
)

func main() {
	// var x int = 10
	x := "10"
	fmt.Println(x)

	if x == strconv.Itoa(5) {
		fmt.Println("correct")
	} else if x == "10" {
		fmt.Println("this is string")
	} else {
		fmt.Println("not correct")
	}

	// if 8%4 == 0 {
	// 	fmt.Println("8 is divisible by 4")
	// }
}
