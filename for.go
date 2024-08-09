package main

import "fmt"

func main() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 0; j <= 5; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("loop")
		break
		// this line will not executed
		fmt.Println("loop -n")
	}

	for i := range []interface{}{"irwan", 1, 2} {
		fmt.Println("range", i)
	}

	for i := range []int{0, 1, 2} {
		fmt.Println("range", i)
	}
}
