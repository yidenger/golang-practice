package main

import "fmt"

func main() {
	s := make([]string, 3)
	fmt.Println("emp: ", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("set: ", s, len(s), cap(s))

	s = append(s, "d")
	s = append(s, "e")

	fmt.Println("apd: ", s, len(s), cap(s))

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy: ", c)

	fmt.Println("s[2:4]: ", s[2:4])

	l := s[1:]
	fmt.Println("l: ", l)
}