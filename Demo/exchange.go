package main

import "fmt"

func main() {
	a := 3;
	b := 6;
	fmt.Println("before exchange: ", a, b);
	a, b = b, a;
	fmt.Println("after exchange: ", a, b);
}