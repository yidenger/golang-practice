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

	var arr3 [2]int = [2]int{21, 22}
	var arr4 [2]int

	arr4 = arr3

	fmt.Printf("arr3: %p, %v\n", &arr3, arr3)
	fmt.Printf("arr4: %p, %v\n", &arr4, arr4)

	fmt.Println(arr3 == arr4)

	var arr
}