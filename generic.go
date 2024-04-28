package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type node[T any] struct{
	next *node[T]
	val T
}

func findIndex[T comparable](arr []T, val T) int{
	for i,v:=range arr{
		if(v==val){
			return i;
		}
	}
	return -1;
}

func isGreater[T constraints.Ordered](x,y T) bool{
	return x>y;
}


func main(){
	fmt.Println("Hola");
	fmt.Println(isGreater(2.3,4.5));
}