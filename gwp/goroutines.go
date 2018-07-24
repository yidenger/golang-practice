package main

import (
	"fmt"
	"time"
)

func printNumbers1() {
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}

func printLetters1() {
	for i := 'A'; i < 'A' + 10; i++ {
		time.Sleep(1 * time.Millisecond)
		fmt.Print("%c ", i)
	}
}

func print1() {
	printNumbers1();
	printLetters1();
}

func goPrint1() {
	go printNumbers1()
	go printLetters1()
}

func main() {
	goPrint1()
	time.Sleep(30 * time.Millisecond)
}