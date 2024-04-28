package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Time Study")

	presentTime := time.Now()
	fmt.Println(presentTime)

	//Formating
	formatedTime1 := presentTime.Format("01-02-2006 15:04:05 Monday MST")
	fmt.Println(formatedTime1)

	newDate := time.Date(2004,time.September,22,4,4,4,4,time.UTC)
	fmt.Println(newDate)


}
