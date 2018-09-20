package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["age"] = 22
	m["heigh"] = 177
	m["weight"] = 140
	
	fmt.Println("map: ", m, len(m))

	delete(m, "heigh")

	fmt.Println(m)

	x, y := m["age"]
	fmt.Println(x, y)

	for key, value := range m {
		fmt.Println(key, value);
	}

	type People struct {
		name string
		id int
	}

	p := make(map[People]int)

	p[People{name: "Jack", id: 1}] = 101
	p[People{name: "Anne", id: 2}] = 102

	fmt.Println("p: ", p)
}