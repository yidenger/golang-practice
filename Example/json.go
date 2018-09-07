package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	FirstName string
	LastName string
}

func main() {
	theJson := `{"FirstName": "Bob", "LastName": "Simth"}`
	var person Person

	json.Unmarshal([]byte(theJson), &person)

	fmt.Printf("%+v\n", person)

	people := Person{"James", "Bond"}

	peopleJson, _ := json.Marshal(people)
	fmt.Printf("%+v\n", string(peopleJson))
}