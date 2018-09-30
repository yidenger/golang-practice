package main

import "fmt"
import "time"
import "sync"

func main() {
	var x int = 0

	for i := 0; i < 2; i++ {
		go func() {
			sync.Mutex.Lock()
			x = x + 1
			sync.Mutex.Unlock()
		}()
	}

	time.Sleep(time.Millisecond)
	fmt.Println("x: ", x);
}