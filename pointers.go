package main
import "fmt"

type points struct{
	x float32
	y float32
}

func main(){
	fmt.Println("Go pointers concept")
	myNumber := 20

	var x *int
	fmt.Println("Default value of x is ",x)

	x = &myNumber
	fmt.Println("New address of x is ",x)
	fmt.Println("New value of x is ",*x)
	
	*x = *x + 19
	fmt.Println("x = ",*x)
	fmt.Println("myNumber = ",myNumber)

	//Pointer in struct
	fmt.Println("Pointers in struct")
	A:=points{
		x:168.92,y: 123.96,
	}

	p := &A;
	fmt.Println(p.x)

}