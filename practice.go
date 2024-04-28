package main

import (
	"fmt"
	"math"
)

// fibonacci is a function that returns
// a function that returns an int.
type Vertex struct {
	x,y int;
}

func (v Vertex) calDistance() float64{
	return math.Sqrt(float64((v.x)*(v.x)+(v.y)*(v.y)))
}


func main() {
	v := Vertex{1,2};
	fmt.Println(v.calDistance())
}
