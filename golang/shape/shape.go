package main

import "fmt"
import "math"

func main() {

	rect := Rectangle{15, 20}
	circ := Circle{15}

	fmt.Println("Rectangle area:", rect.area())
	fmt.Println("Circle area:", circ.area())

}

type Shape interface {
	area() float64
}

type Rectangle struct {
	height float64
	width  float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) area() float64 {
	return r.height * r.width
}

func (c Circle) area() float64 {
	return math.Pow(c.radius, 2) * math.Pi
}

func getArea(shape Shape) float64 {
	return shape.area()
}
