package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}
type rectangle struct {
	width  float64
	height float64
}

type circle struct {
	radius float64
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func getArea(s shape) float64 {
	return s.area()

}

func main() {
	fmt.Println("Hello World")
	r1 := rectangle{height: 10, width: 20}
	c1 := circle{radius: 5}
	shapes := []shape{c1, r1}

	for _, shape := range shapes {
		fmt.Println(getArea(shape))
	}

}
