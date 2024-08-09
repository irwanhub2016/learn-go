package main

import "fmt"

func main() {
	const pi = 3.14
	var radius float64

	fmt.Println("Please input radius: ")
	fmt.Scanln(&radius)

	area := pi * radius * radius

	fmt.Println(area)
}
