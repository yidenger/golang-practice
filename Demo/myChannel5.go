package main

import "fmt"

func main() {
	number := 589

	sqrch := make(chan int)
	cubech := make(chan int)
	go calcCubes(number, sqrch)
	go CalcSquares(number, cubech)

	squares, cubes := <-sqrch, <-cubech

	fmt.Println("Final output: ", squares + cubes)
}

func CalcSquares(number int, squareop chan int) {
	sum := 0
	
	for number !=0 {
		digit := number % 10
		sum += digit * digit
		number /= 10
	}

	squareop <- sum
}

func calcCubes(number int, cubeop chan int) {
	sum := 0

	for number != 0 {
		digit := number % 10
		sum += digit * digit * digit
		number /= 10
	}

	cubeop <- sum
}