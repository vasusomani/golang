package main

import (
	"errors"
	"fmt"
)

func throwError(x int) (int, error) {
	if x > 10 {
		//return x,fmt.Errorf("Error! %v","Value is greater than 10")
		return x, errors.New("Error! value is greater than 10")
	}
	return x, nil

}

func main() {
	_, error := throwError(15)
	fmt.Println(error)
}
