package main

import "fmt"

func main() {
	months := [...]string{1: "A1", 2: "A2", 3: "A3", 4: "A4", 5: "A5", 6: "A6", 7: "A7", 8: "A8", 9: "A9", 10: "A10", 11: "A11", 12: "A12"};
	fmt.Println(months[1:13], cap(months));
}