package main

import "fmt"

func MultiReturn() (int, int) {
	return 5, 6

}

func main() {
	a, b := MultiReturn()
	fmt.Println(a)
	fmt.Println(b)

	_, c := MultiReturn()
	fmt.Println(c)
}
