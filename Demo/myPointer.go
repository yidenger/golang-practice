package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
} 

func main() {
	a := 1
	b := 2

	zeroval(a)

	fmt.Println(a)

	zeroptr(&b)
	fmt.Println(b)
}