package main

import "fmt"

func main() {
	s := make([]int, 5)

	fmt.Printf("%p\n", s)
	s = append(s, 1, 2, 3)
	fmt.Printf("%p\n", s)
	fmt.Println(s)
}

// address
// new address
// 0 0 0 0 0 1 2 3