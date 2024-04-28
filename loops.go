package main
import "fmt"

func main(){
	n:=5
	//For loops dont use any circular parenthesis
	//For Loop
	for i:=0;i<n;i++{
		fmt.Print(i," ")
	}
	fmt.Println()

	//There is no While Loop in go,instead use for along with condition
	i:=0
	for ;i<n;{
		fmt.Print(i," ")
		if(i==2){
			break
		}
		i++
	}
	//OR
	for i<n{
		fmt.Print(i," ")
		if(i==2){
			break
		}
		i++
	}
	fmt.Println()

	//infinite loop
	// for{
		
	// }

	//For each loop
	charSlice := []string{"A","B","C","D","E"}
	for index:=range charSlice{
		fmt.Println(charSlice[index]," ")
	}
	
	for index,value:=range charSlice{
		fmt.Printf("index : %v, value: %v\n",index,value)

	}
	for _,v := range charSlice{
		fmt.Println(v)
	}
	
	goto mark1//goto Statement

	mark1:	//goto label 
		fmt.Println("Its Mark 1")


}
