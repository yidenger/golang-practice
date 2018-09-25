package main

import "fmt"
import "time"

func main() {
	Publish("A goroutine starts a new thread of execution.", 5 * time.Second)
	fmt.Println("Let's hope the news will published before I leave")

	time.Sleep(10 * time.Second)
	fmt.Println("Ten seconds later: I'm leaving now.")
}

func Publish(text string, delay time.Duration) {
	go func() {
		time.Sleep(delay)
		fmt.Println("BREAKING NEWS: ", text)
	}()
}