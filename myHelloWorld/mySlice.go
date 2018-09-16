package main

import "fmt"

func mySlice() {
	slice1 := make([]int, 5);
	fmt.Println("slice1: ", slice1, len(slice1));

	slice2 := make([]int, 5, 10);
	fmt.Println("slice2: ", slice2, len(slice2), cap(slice2));

	slice3 := []int{1, 2, 3, 4, 5};
	fmt.Println("slice3: ", slice3, len(slice3), cap(slice3));

	arr1 := [5]int{5, 6, 7, 8, 9};

	slice4 := arr1[0:3];
	fmt.Println("slice4: ", slice4, len(slice4), cap(slice4));

	slice5 := []int{10, 11, 12, 13, 14};
	slice6 := slice5[2:5];
	fmt.Println("slice6: ", slice6, len(slice6), cap(slice6));

	slice6 = append(slice6, 15);
	fmt.Println("slice6: ", slice6, len(slice6), cap(slice6));

	slice6 = append(slice6, 15);
	fmt.Println("slice6: ", slice6, len(slice6), cap(slice6));

	slice7 := []int{20, 21, 22, 23, 24, 25}
	slice8 := slice7[1: 4];
	fmt.Println("slice8: ", slice8);

	slice9 := make([]int, len(slice8));
	copy(slice9, slice8);

	slice8[0] = 99;
	fmt.Println("after slice8: ", slice8);
	fmt.Println("after slice7: ", slice7);
	fmt.Println("after slice9: ", slice9);

}

func main() {
	fmt.Println("hello my slice");
	mySlice();
}