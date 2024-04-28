package main

import (
	"fmt"
)

// Interface group types together based on their methods
type shape interface {
	area() float64
	perimeter() float64
}

//Only for those structures having both area and perimeter methods with float64 as their return type

type rectangle struct { //In java to make it a part of interface we do -> type rectangle struct implements shapes{}..... but we do directly in go
	length  float64
	breadth float64
}

type square struct {
	size float64
}

// Area Methods
func (r rectangle) area() float64 {
	return r.length * r.breadth
}

func (s square) area() float64 {
	return s.size * s.size
}

// Perimeter Methods
func (r rectangle) perimeter() float64 {
	return 2 * (r.length + r.breadth)
}

func (s square) perimeter() float64 {
	return 4 * s.size
}

// Using Interface as parameter=> Here s shapes represents all the structs inside the interface
func printArea(s shape) {
	fmt.Println(s.area())
	fmt.Println(s.perimeter())
	fmt.Println()
}

// Typecasting an interface
func getSquare(s shape) {
	sq, ok := s.(square) //sq is square casted from shape and ok is boolean denoting whether it is casted succesfully or not. If it fails(when the data is not from square struct) then sq is nil and ok is false
	fmt.Println(sq, ok)
}

// Type switch
func interfaceTypeSwitch(s shape) {
	switch s.(type) {
	case rectangle:
		fmt.Println("Rectangle")
	case square:
		fmt.Println("Square")
	default:
		fmt.Println("Pta ni bhai")
	}
}

func main() {
	r := rectangle{length: 10, breadth: 2}
	s := square{size: 5}
	// c:=circle{radius:3}
	printArea(r)
	printArea(s)
	interfaceTypeSwitch(r)
}
