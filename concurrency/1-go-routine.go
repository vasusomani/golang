package main

import (
	"fmt"
	"time"
)

//Concurrency is a concept of executing multiple tasks at the same time
//by utilizing all available resources more effectively

func attack(agent string){
	fmt.Printf("%v attack\n",agent)
	time.Sleep(time.Second);
}

func aamZindagi(agents []string){
	start := time.Now();
	defer func(){
		fmt.Println(time.Since(start));
	}();

	for _,v := range agents{
		attack(v); //Aam Zindagi (Slow)
	}
}

func mentosZindagi(agents []string){
	start := time.Now();
	defer func(){
		fmt.Println(time.Since(start));
	}();

	for _,v := range agents{
		go attack(v); //Goroutine	 //Mentos Zindagi (Fast)
		// Goroutines run in the same address space, so access to shared memory must be synchronized.
	}
}

func main(){
	agents := []string{"Agent 1","Agent 2","Agent 3"};
	aamZindagi(agents)
	mentosZindagi(agents)
	time.Sleep(time.Second) //since we are not using go channel so sometimes main function executes and end before our routine function so we are giving a manual appx delay to main fxn
}

/* 
// Without go-channel
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
	 |										|
	 |										|
	 |										|
	 |										|
	 |(main fxn completed)					| (task fxn still not completed)
	 |(manual time delay)					|
	 |(manual time delay)					|
	 |(manual time delay)					|
	 |(manual time delay) <----------------	| (task fxn completed)
	 |(manual time delay)					
	 |(manual time delay)					
	 |(delay over) (main fxn ends)		
	 

	 //Since we are not using go channel so sometimes main function executes and end before our routine function
	 //So we will give manual delay to main fxn so that it runs for more time and won't finish up before task fxn
	 //Hence we won't get proper results i.e(won't end main fxn just after executing task fxn, might get additional delay)
*/