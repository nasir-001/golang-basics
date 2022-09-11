package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
}

type Circle struct {
	x, y, radius float64
}

type Rectangle struct {
	width, height float64
}

func (self Circle) area() float64 {
	return math.Pi * self.radius * self.radius
}

func (self Rectangle) area() float64 {
	return self.width * self.height
}

func getArea(self Shape) float64 {
	return self.area()
}

func main() {
	circle := Circle{x: 0, y: 0, radius: 5}

	rectangle := Rectangle{width: 10, height: 5}

	fmt.Printf("Circle Area: %f\n", getArea(circle))
	fmt.Printf("Rectangle Area: %f\n", getArea(rectangle))
}
