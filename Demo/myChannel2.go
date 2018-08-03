package main

import "fmt"

func main() {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	fmt.Println("222")
	ch <- 3
	fmt.Println("333")
	fmt.Println(<-ch)
	ch <- 4
	fmt.Println("444")
	for v := range ch {
		fmt.Println(v)
		if len(ch) <= 0 {
			break
		}
	}
}