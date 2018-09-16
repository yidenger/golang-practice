package main

import "fmt"

func myMap() {

	map1 := make(map[string]int);
	map1["age"] = 22;
	fmt.Println(map1);

	map2 := map[string]int{"age": 18}
	fmt.Println(map2);

	map3 := map[string]int{"age": 18, "w": 11, "h": 8}

	for k, v := range map3 {
		fmt.Printf("%s: %d", k, v);
		fmt.Println();
	}

}

func main() {
	fmt.Println("hello my map");

	myMap();
}