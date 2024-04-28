//Golang is Compiled lang : Go tool can run file directly,without any VM
//Not OOP, Operator Overloading, try-catch

package main
//All the methods from any package are named with a capital letter (eg: fmt.Println()). It means this function is exported publicly

import (
	"fmt"
	"runtime"
	"time"
)

//Global Variable
var name1 string = "Vasu" //allowed
//name2 := "Vasu" //Not allowed to use walrus operator outside any method

func main() {
	startTime := time.Now()
	MAIN()
	fmt.Println("---------------------------------\n")
	STRUCT()
	endTime := time.Now()
	fmt.Println("\n\n---------------------------------------------------------------------------")
	// Calculate and print the execution time
	executionTime := endTime.Sub(startTime)
	fmt.Println("Execution Time:", executionTime)

	// Get memory usage statistics
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Allocated Memory: %v bytes\n", m.Alloc)
	fmt.Printf("Total Memory Allocated (including freed): %v bytes\n", m.TotalAlloc)
	fmt.Printf("System Memory: %v bytes\n", m.Sys)
	fmt.Printf("Number of Garbage Collections: %v\n", m.NumGC)

}



func MAIN() {
	fmt.Println("Hello World")
	// Declaration Type-1
	var x int = 10
	fmt.Println(x)
	// Declaration Type-2 (No need to give dtype name it will automatically assign it)
	// Type 2 is only valid for declaring a var inside a fxn, it wont work outside a fxn
	y := 20
	z := 0.2
	//OR y,z := 20,0.2
	var a float32 = 9.7
	b := true
	//Public variable
	Name := "Vasu" //Can be accessed from anywhere (Make first letter captial), similar to public keyword in Java
	fmt.Println(Name)
	//Multiple prints using comma
	fmt.Println(y, x)

	//Format Specifier
	fmt.Printf("Just cheking for type : %T", y)
	fmt.Print("\n") //similar to fmt.Println()
	fmt.Printf("Just cheking for type  : %d \n", y)
	fmt.Printf("Just cheking for type : %T\n", z)
	fmt.Printf("Just cheking for type : %T\n", b)
	fmt.Printf("Just cheking for type : %v\n", a) //Automatic specifier

	//Constants :- You cant change its value (== final in Flutter)
	//cannot be declared using ":="
	const name = "Vasu"
	//name="xyz" WRONG
	fmt.Println(name)

	//sprintf : allows you to create formatted strings by substituting values into placeholders within the string.
	msg := fmt.Sprintf("Hello Guys my age is %v", "Vasu")
	fmt.Println(msg)

	data(1,5)

	//multiple returns in a fucntion
	r1,r2 := getPoint()
	fmt.Printf("%v %v \n",r1,r2)
	r3,_ := getPoint() //ignore 2nd return using '_'
	fmt.Printf("%v \n",r3)

	//type conversion
	i := 20
	var ith float32 = float32(i)
	fmt.Printf("%T %v \n",ith,ith)
}
	


//Function
//Varibale name followed by datatype
//For void functions
func data(x int,y int){
	fmt.Println(x+y);
}

//For non-void functions
func sub(x int,y int) int{
	return x-y
} 

func add(x, y int) int{
	return x+y
} 

//Multiple indefenite inputs
func adder(values... int) int{
	sum:=0
	for _,val := range values{
		sum+=val
	}
	return sum
}

//Complex function declaration
func fr(a func(int,int) int,b int) int{
	return 1
}

//Multiple returns
func getPoint() (x, y int) {
	return 3, 4
}

//Naked returns
func getPoint1() (x,y int){
	//x and y are auto initialzied to 0

	return //automatically returns x,y in order => Similar to return x,y
}


func STRUCT(){
	//Declaration
	type model struct{
		modelName string
		modelColor string
	}
	//Nested struct
	type car struct{
		name string
		age int
		carModel model
		number int
	}
	//Anonymous Struct : No naming is done
	type bike struct{
		name string
		age int
		bikeModel struct{
			modelName string
			modelColor string
		}
		number int
	}
	//Embeded Struct
	type truck struct{
		name string
		age int
		model //Inherited from model struct
		number int
	}

	alto := car{name: "Alto",age: 5,carModel: model{modelName: "Alto800",modelColor: "Red"}}
	alto.number = 12345
	fmt.Println(alto.carModel.modelName)
	//Using embeded struct
	tata := truck{name: "Tata",age: 12,number: 696969,model: model{modelName: "TataRTX",modelColor: "White"}}
	tata.modelColor="Red"
	fmt.Println(tata.modelColor) //Instead of tata.model.modelColor in nested struct
	fmt.Printf("My Car details : %v",tata)//only values
	fmt.Println()
	fmt.Printf("My Car details : %+v",tata)//with key and values -> only for struct %+v
}

//INTERFACE
type rect struct {
	width, height float64
}

func(r rect) area() float64{
	return r.height*r.width
}

