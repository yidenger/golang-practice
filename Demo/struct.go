package main

import (
	"fmt"
)

func main() {
	type Employee struct {
		ID int 
		Name string
		Address string
		Position string
		Salary int
		ManagerID int
	}

	var dilbert Employee

	dilbert.Salary = 5000;

	fmt.Println(dilbert);
}