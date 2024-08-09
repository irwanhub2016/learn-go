package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	for i := 0; i < 1000; i++ {
		fmt.Println("Hello, World!")
	}

	elapsed := time.Since(start)
	fmt.Printf("Go took %s\n", elapsed)
}
