package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your rating: ")
	//Comma ok syntax
	rating, _ := reader.ReadString('\n') //Read till \n is encountered

	//Type conversions
	// rating_int, error := strconv.ParseInt(strings.Trim(rating,"\n"), 10, 32)
	rating_int, error := strconv.ParseInt(strings.TrimSpace(rating), 10, 32)
	if error != nil {
		fmt.Println(error)
	} else {
		rating_int += 1
		fmt.Println("New rating is :", rating_int)
	}
}

//Memorary
// Memory allocation and deallocation occurs automatically
// new() to allocate memory but not initialize
// make() to allocate memory and initialize
//Garbage collection occurs automatically