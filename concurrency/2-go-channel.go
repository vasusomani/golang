package main

import (
	"fmt"
	"time"
)

//Concurrency is a concept of executing multiple tasks at the same time
//by utilizing all available resources more effectively

func attack(agent string, isDone chan bool){
	fmt.Printf("%v attack\n",agent)
	time.Sleep(time.Second);
	isDone <- true; //For assigning value to channel use "channel"<-
}

func main(){
	agents := []string{"Agent 1","Agent 2","Agent 3"};
	start := time.Now();
	defer func(){
		fmt.Println(time.Since(start));
	}();

	isDone := make(chan bool)
	// isDone <- false 	//Deadlock Error : Since default capacity of a channel is 0(i.e no buffer)
						//If I send some value to channel from main go-routine side
						//Then other side there is no-one who will take or consume that value hence its a deadlock

	for _,v := range agents{
		go attack(v,isDone);
		fmt.Println(<-isDone); //For extracting value from channel use <-"channel"
	}
}


/* 
// With go-channel
	main							task(go-routine)
	 |										
	 |
	 |
	 |
	 |		Called go-routine function
	 | -----------------------------------> |
	 |										|
	 |										|
	 |										|
	 |		   *Shared Memory (SM)*			|
	 |		   *Periodically Check* 		|
	 |										|
	 |							  			|
	 |(main fxn completed)					| (task fxn still not completed)
	 |										|
	 |										|
	 |										|
	 | <----------------------------------- | (task fxn completed)									
			main fxn also ends here


	 //When we are done with task process we add some check mark to shared memory
	 //Main fxn will be periodically checking SM and when it gets the signal that task process is done
	 //Then it will end the main go-routine or main fxn
	 //The go-channel acts as the shared memory here
*/