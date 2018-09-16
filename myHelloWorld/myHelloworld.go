package main

import "fmt"

func myArr() {
	var arr1 [5]int;

	fmt.Println("arr1: ", arr1, len(arr1), cap(arr1));

	for i :=0; i < 5; i++ {
		arr1[i] = i + 1;
	}

	fmt.Println("arr1: ", arr1, len(arr1), cap(arr1));

	arr2 := [3]int{5, 6, 7};
	fmt.Println("arr2: ", arr2, len(arr2));

	arr3 := [10]int{0};
	fmt.Println("arr3: ", arr3, len(arr3));

	arr4 := [...]int{10, 11, 12, 13, 14, 15, 16};
	fmt.Println("arr4: ", arr4, len(arr4));
	
	arr5 := [5]int{1:9};
	fmt.Println("arr5: ", arr5, len(arr5));

	var twoD [2][3]int;

	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j;
		}
	}

	fmt.Println("twoD: ", twoD);
}

func main() {
	fmt.Println("hello world");
	myArr();
}
