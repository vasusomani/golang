package main

import (
	"fmt"
	"math/rand"
	"time"
)

func shoot(channel chan string){
	start:= time.Now();
	rounds:=5
	for i:=0;i<rounds;i++{
		time.Sleep(time.Second) ;
		channel <- fmt.Sprintf("Your score is %d",rand.Intn(10));
	}
	close(channel); //without it, we will get deadlock since we have extracted all the messages but channel is still open
	//Note: Only the sender should close a channel, never the receiver.
	// Sending on a closed channel will cause a panic.
	//Channels aren't like files; you don't usually need to close them
	//Closing is only necessary when the receiver must be told there are no more values coming, such as to terminate a range loop.
	fmt.Println(time.Since(start));
}

func main(){
	channel := make(chan string);
	go shoot(channel);

	// for message := range channel{
	// 	fmt.Println(message)
	// }

	//or

	for {
		message,open := <- channel
		if(!open){
			break;
		}
		fmt.Println(message);
		
	}
}