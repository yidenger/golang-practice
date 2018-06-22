package main

import "fmt"

func main() {
	str := "hi,你好world"
	fmt.Println("length: ", len(str));
	fmt.Println(str[2], str[4], str[6]);
}