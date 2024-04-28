package main

import (
	"fmt"
	"time"
)

//select statement lets a goroutine wait on multiple communication operations

func goElection(channel chan string,ninja string){
	time.Sleep(time.Second)
	channel <- ninja;
}

func roughlyFair(){
	heads := make(chan interface{});close(heads);
	tails := make(chan interface{});close(tails);

	hCount,tCount := 0,0;

	for i:=0;i<500;i++{
		select{
		case <-heads:
			hCount++;
		case <-tails:
			tCount++;
		}
	}
	fmt.Println("Heads =",hCount,"Tails count =",tCount);

}

func main(){
	//Example 1
	ninja1,ninja2 := make(chan string),make(chan string);

	go goElection(ninja1,"Ninja 1");
	go goElection(ninja2,"Ninja 2");
	
	select{ //It will randomly select between any of its cases
	case message := <-ninja1:
		fmt.Println(message);
	case message := <-ninja2:
		fmt.Println(message);
	// default:
	// 	fmt.Println("Default"); //Default takes min time since no go-routine is associated with it so always default will come
	}

	//Example 2
	roughlyFair()

}