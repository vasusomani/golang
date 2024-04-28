package main

import (
	"fmt"
	"sync"
	"time"
)

func completeTask(task string, wg *sync.WaitGroup){
	fmt.Println(task,"completed");
	time.Sleep(time.Second);
	wg.Done();
}

func main(){
	var wg sync.WaitGroup;
	tasks := []string{"Task 1","Task 2","Task 3","Task 4"};
	wg.Add(len(tasks));

	for _,task := range tasks{
		go completeTask(task,&wg);
	}
	
	wg.Wait(); // wait till the sub-routine completes
	//The sequence of sub-routines may be different but till all sub-routines are executed the wg will will
	fmt.Println("All tasks completed");
}