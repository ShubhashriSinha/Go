package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
	circumf() float64
}

type square struct {
	length float64
}
type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.length * s.length
}
func (s square) circumf() float64 {
	return s.length * 4
}

func (c circle) area() float64 {
	return c.radius * c.radius * math.Pi
}
func (c circle) circumf() float64 {
	return c.radius * 2 * math.Pi
}

func printShapeInfo(s shape) {
	fmt.Printf("Area of %T is: %0.2f \n", s, s.area())
	fmt.Printf("Circumference of %T is: %0.2f \n", s, s.circumf())
}

func main() {
	shapes := []shape{
		square{length: 10},
		circle{radius: 3},
		circle{radius: 2},
		square{length: 8},
	}

	for _, v := range shapes {
		printShapeInfo(v)
		fmt.Println("---")
	}

}
