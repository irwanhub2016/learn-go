package main

import "fmt"

func hitung(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func main() {
	ayoHitung := hitung(5, 5)
	fmt.Println(ayoHitung)

	res := plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)
}
