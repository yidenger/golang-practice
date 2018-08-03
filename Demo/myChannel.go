package main 

import "fmt"

func main() {
	messages := make(chan string)

	go func () { 
		fmt.Println("do something")
		// messages <- "ping"
	}()

	msg := <- messages
	fmt.Println(msg)
}