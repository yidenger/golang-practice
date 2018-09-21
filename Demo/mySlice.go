package main

import "fmt"

func main() {
	var arr [4]int = [4]int{1, 2, 3}

	fmt.Println(arr)

	b := [...]string{"penn", "teller"}
	fmt.Println(b, len(b))

	var s []int = []int{10, 11, 12, 13, 14}
	fmt.Println(s, len(s), cap(s))
	fmt.Println(s[1:4])

	arr2 := arr
	arr2[1] = 100

	fmt.Println(arr, arr2)
}