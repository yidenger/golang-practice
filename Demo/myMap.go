package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["age"] = 22
	m["heigh"] = 177
	
	fmt.Println("map: ", m, len(m))

	delete(m, "heigh")

	fmt.Println(m)

	x, y := m["age"]
	fmt.Println(x, y)
}