package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4}
	sum := 0
	
	for index, num := range nums {
		fmt.Println("index: ", index)
		sum += num
	}

	fmt.Println("sum: ", sum)

	person := map[string]string {
		"age": "22",
		"name": "jack",
		"hobby": "singing",
	}

	for k, v := range person {
		fmt.Printf("%s => %s\n", k, v)
	}
}