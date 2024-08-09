package main

import (
	"fmt"
	"math"
)

func main() {
	const score = 12345
	fmt.Println(score)

	const n = 100000

	// const score = 598334
	// fmt.Println(score)

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))
	fmt.Println(math.Sin(n))
}
