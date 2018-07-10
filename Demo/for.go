package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}

	fmt.Println("sum: ", sum)

	arr := [...]string{"xiaoming", "liuxiang", "yaoming", "liuyue"}

	for key, value := range arr {
		fmt.Println("key, value: ", key, value)
	}
}