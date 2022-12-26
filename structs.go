package main

import (
	"fmt"
	"math"
)

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

type Circle struct {
	x, y, r float64
}

//	func area (c Circle) float64 {
//		return c.r * c.r * math.Pi
//	}
// fmt.Println(circleArea(c))

//	func area (c *Circle) float64 {
//		return c.r * c.r * math.Pi
//	}
// fmt.Println(circleArea(&c))

func (c *Circle) area() float64 {
	return c.r * c.r * math.Pi
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

func main() {
	//var c Circle
	//c := new(Circle)
	c := Circle{2, 3, 1}
	fmt.Println("Circle:", c.area())

	r := Rectangle{0, 0, 10, 10}
	fmt.Println("Rectangle:", r.area())

}
