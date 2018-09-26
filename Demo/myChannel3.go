package main

import (
	"sync"
	"time"
	"fmt"
)

func main() {
	// ch := make(chan string)

	// go func() {
	// 	ch <- "Hello!"
	// 	close(ch)
	// }()

	// fmt.Println(<-ch)
	// fmt.Println(<-ch)
	// fmt.Println(<-ch)

	// v, ok := <-ch

	// fmt.Println(v, ok)


	// var ch2 <-chan string = Producer()
	// for s := range ch2 {
	// 	fmt.Println("Consumed: ", s)
	// }


	// wait := Publish("A goroutine starts a new thread of execution.", 5 * time.Second)
	// fmt.Println("Let's hope the news will published before I leave")
	// <-wait
	// fmt.Println("Ten seconds later: I'm leaving now.")

	// race()

	correct()
}

func Producer() <-chan string {
	ch := make(chan string)
	go func() {
		ch <- "Hi"
		ch <- "Xiaoming"
		ch <- "Nihao"
		close(ch)
	}()

	return ch
}

func Publish(text string, delay time.Duration) (chan string){
	ch := make (chan string)
	go func() {
		time.Sleep(delay)
		fmt.Println("Breaking news: ", text)
		ch <- "nihao"
		// close(ch)
	}()
	return ch
}

func race() {
	ch := make(chan int)

	
	go func() {
		n := 0
		n++
		fmt.Println("xx: ", n)
		ch <- n
	}()
	
	n := <-ch
	n++
	fmt.Println("yy: ", n)
	
}

func correct() {
	var wg sync.WaitGroup
	wg.Add(5)

	for i := 0; i < 5; i++ {
		go func () {
			fmt.Print(i)
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println()
}