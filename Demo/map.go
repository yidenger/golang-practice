package main

import "fmt"

func main() {
	person := map[string]int{
		"age": 22,
		"height": 178,
	}

	fmt.Println(person);
	fmt.Println(person["age"]);
	fmt.Println(person["height"]);
	delete(person, "height");
  person["width"] = person["width"] + 1;
	fmt.Println(person);

	for key, value := range person {
		fmt.Printf("key: %s\t value: %d\n", key, value);
	}

	tall, ok := person["tall"]
	fmt.Println("ok: ", ok, tall);
	
	age, ok := person["age"]
	fmt.Println("ok: ", ok, age);
}