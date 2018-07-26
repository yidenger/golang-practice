package main

import "fmt"
import "math"

type geometry interface {
	area() float64
	perim() float64
}

type react struct {
	width, height float64
}

type circle struct {
	radius float64
}