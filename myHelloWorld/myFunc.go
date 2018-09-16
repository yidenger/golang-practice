package main

import "fmt"

func sum(a int, b int) int {
	return a + b;
}

func multi(name string, age int) (string, int){
	return name, age;
}

func main() {
	fmt.Println("hello my function");

	fmt.Println("sum(1, 2): ", sum(1, 2));

	name, age := multi("xiaoming", 18);

	fmt.Printf("my name is %s, i am %d years old", name, age);
}