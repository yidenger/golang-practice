package main

import "fmt"

func main() {
	number := 5

	switch {
		case number > 2:
			fmt.Println("> 2")
		case number > 4:
			fmt.Println("> 4")
		case number > 5:
			fmt.Println("> 5")
		case number > 8:
			fmt.Println("> 8")
	}

}