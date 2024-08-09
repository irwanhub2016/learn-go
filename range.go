package main

import "fmt"

func main() {
	coffe_menu := []string{"espresso", "latte", "americano"}

	for i, menu := range coffe_menu {
		fmt.Println(i, menu)
	}
}
