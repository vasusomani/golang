package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){
	rand.Seed(time.Now().UnixMicro())
	randomNumber := rand.Intn(6)+1
	switch randomNumber{
	case 1 : fmt.Println("Its 1")
	case 2 : fmt.Println("Its 2")
	case 3 : fmt.Println("Its 3")
			fallthrough
	case 4 : fmt.Println("Its 4")
	case 5 : fmt.Println("Its 5")
	case 6 : fmt.Println("Its 6")
	}
	//In go no need to use break statement after every case.
	//But if you want to go to next case if one opens then you can use "fallthrough"
}