package main

import "fmt"

type Human struct {
	name string
	age int
	weight int
}

type Student struct {
	Human
	speciality string
	age int
}

func main() {
	mark := Student{Human{"Mark", 25, 140}, "Computer Science", 22}

	fmt.Println(mark)
	fmt.Println(mark.age)
	fmt.Println(mark.Human.age)
}