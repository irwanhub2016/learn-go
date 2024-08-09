package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["one"] = 1
	m["two"] = 2
	m["three"] = 3

	fmt.Println(m)

	v1 := m["one"]
	fmt.Println(v1)

	delete(m, "three")
	fmt.Println(m)

	// clear(m)
	// fmt.Println(m)

	// return boolean
	_, prs := m["one"]
	fmt.Println("prs:", prs)

	n := map[string]int{"foo": 1, "took": 2, "cops": 3}

	value, ok := n["foo"]
	if ok {
		fmt.Println("exist", value)
	} else {
		fmt.Println("not exist")
	}

	for key, value := range n {
		fmt.Println(key, value)
	}
}
