package main

import "fmt"

type person struct {
	name string
	age int
}

func main() {
	fmt.Println(person{"Bob", 22})

	fmt.Println(person{name: "xiaoming", age: 24})

	fmt.Println(person{name: "Jack"})

	fmt.Println(person{age: 20})

	p := person{
		name: "Lucy",
		age: 23}

	sp := p

	sp.age = 27

	fmt.Println(p, sp)
	fmt.Println(sp)
}