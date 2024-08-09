package main

import (
	"fmt"
)

func main() {
	var arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var new_arr []int

	// filter number
	for j := 0; j <= len(arr)-1; j++ {
		if arr[j]%3 == 0 || arr[j]%5 == 0 {
			new_arr = append(new_arr, arr[j])
		}
	}

	// calculate result
	var total = 0
	for _, numb := range new_arr {
		total += numb
	}
	
	fmt.Println(total)
}
