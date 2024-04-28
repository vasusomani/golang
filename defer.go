package main
import "fmt"

//defer function or method call arguments evaluate instantly
//but they don’t execute until the nearby functions returns
//If multiple defer statements are there then LIFO is followed
func main(){
	fmt.Println("First")
	defer fmt.Println("Second")
	defer fmt.Println("Third")
	fmt.Println("Fourth")
	defer delayedFunc()
}

func delayedFunc(){
	fmt.Println("To be delayed")
}