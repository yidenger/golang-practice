package main

import "fmt"

func main() {
	defer_call()
}

func defer_call() {
	defer func() {
		fmt.Println("打印前") // 1
	}()

	defer func() {
		fmt.Println("打印中") // 2
	}()

	defer func() {
		fmt.Println("打印后") // 3
	}()

	panic("触发异常") // 4
}

// 3214