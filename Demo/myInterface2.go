package main

import (
	"fmt"
)

type I interface {
	Get() int
	Set(int)
}

type S struct {
	Age int
}

func (s S) Get() int {
	return s.Age
}

func (s *S) Set(age int) {
	s.Age = age
}

func f(i I) {
	i.Set(10)
	fmt.Println(i.Get())
}

func main() {
	s := S{4}
	f(&s)

	fmt.Println(s)

	var i I
	i = &s
	fmt.Println(i.Get())

	if t, ok := i.(*S); ok {
		fmt.Println("s implements I", t)
	}
}


