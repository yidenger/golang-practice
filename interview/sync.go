package main

import (
	"fmt"
	"sync"
)

type Userages struct {
	ages map[string]int
	sync.RWMutex
}

func (ua *Userages) Add(name string, age int) {
	ua.Lock()
	defer ua.Unlock()
	ua.ages[name] = age
}

func (ua *Userages) Get(name string) int {
	ua.RLock()
	defer ua.RUnlock()

	if age, ok := ua.ages[name]; ok {
		return age
	}

	return -1
}

func main() {
	count := 10

	gw := sync.WaitGroup{}
	gw.Add(count * 3)

	u := Userages{ages: map[string]int{}}

	add := func(i int) {
		u.Add(fmt.Sprintf("user_%d", i), i)
		gw.Done()
	}

	for i := 0; i < count; i++ {
		go add(i)
		go add(i)
	}

	for i := 0; i < count; i++ {
		go func(i int) {
			defer gw.Done()
			u.Get(fmt.Sprintf("user_%d", i))
		}(i)
	}

	gw.Wait()
	fmt.Println("Done")
}