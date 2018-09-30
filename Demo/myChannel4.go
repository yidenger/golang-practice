package main

import "fmt"
import "time"

func main() {
	done := make(chan bool)
	fmt.Println("Main go to call hello go goroutine")
	go hello(done)
	<- done

	fmt.Println("main recevied data")
}

func hello(done chan bool) {
	fmt.Println("Hello go routine is going to sleep")
	time.Sleep(4 * time.Second)
	fmt.Println("hello go routine awake and going to write to done")
	done <- true
	fmt.Println("222")
}
