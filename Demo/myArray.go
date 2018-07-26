package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println("emp: ", a)

	a[4] = 100
	fmt.Println(a, len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("b: ", b, len(b))
}