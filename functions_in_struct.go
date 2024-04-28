package main

import (
	"fmt"
	"math"
)

type rectangle struct {
	length  float64
	breadth float64
}

type circle struct {
	radius float64
}

//struct is passed here by value not refrence
func (r rectangle) area() float64 {
	return r.length * r.breadth
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c *circle) expand(n int){
	c.radius = c.radius*float64(n);
}

func main() {
	r := rectangle{length: 2, breadth: 4}
	c := circle{radius: 2}
	fmt.Println(r.area())
	fmt.Println(c.area())
	c.expand(10);
	fmt.Println(c.radius);
	fmt.Println(c.area());
}
