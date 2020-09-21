package main

import (
	"fmt" 
	"math"
)

func main() {
	circle := Circle{X:2, Y:2 , Radius:10}
	rectangle := Rectangle{width:10, height:5}

	fmt.Println("Area Circle = ", getArea(circle))
	fmt.Println("Area Rectangle = ", getArea(rectangle))
}

type Shape interface {
	Area() float64

}
type Circle struct {
	X   float64
	Y float64
	Radius float64
}

type Rectangle struct {
	width float64
	height float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func getArea(s Shape) float64 {
	return s.Area()
}