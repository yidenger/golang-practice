package main

import (
	"fmt"
	"time"
)

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	go func() {
		fmt.Println("counter: ", time.Now())
		for x := 0; x < 10; x++ {
			naturals <- x
		}
		close(naturals)
	}()

	go func() {
		fmt.Println("squares: ", time.Now())
		for {
			x, ok := <-naturals

			if !ok {
				break;
			}
			squares <- x * x
		}
		close(squares)
	}()

	for {
		value, ok := <- squares
		if !ok {
			break;
		}
		fmt.Println("log: ", value)
	}
}