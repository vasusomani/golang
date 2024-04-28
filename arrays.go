package main

import (
	"fmt"
	"sort"
)

func main(){

	fmt.Println("Arrays")

	var names1 [4]string
	names1[0] = "A"
	names1[1] = "B"
	names1[3] = "D"
	fmt.Println("names1 array : ",names1)

	var numbers1 = [4]int {1,2,3,4}
	// var numbers1 [4]int = [4]int {1,2,3,4}
	fmt.Println("numbers1 array : ",numbers1)

	marks := [4]int{90,95,89,28}
	fmt.Println("marks array : ",marks)
	// marks = append(marks,23)  can't append arrays

	var arr1 [8]int;
	arr1 = [8]int{1,2,3,4,5,6,7,8};
	var arr2 []int;
	arr2 = []int{1,2,3,4,5};
	arr3 := []int{1,2,3,4,5,6,7,8,9,10}

	fmt.Println(arr1[0])
	fmt.Println(arr2[3])
	fmt.Println(arr3[5])

	fmt.Println("Slices")

	subArr1 := arr1[2:5];
	fmt.Println(subArr1[0]);
	//Slicing stores the value address, if value at main arr changes, slicedArr value also changes and vice versa
	arr1[2] = 10;
	fmt.Println(subArr1[0]);
	//A slice has both a length and a capacity.
	//length = no. of elements in the slice => len(arr)
	//capacity = no of elements in the underlying array, counting from the first element in the slice => cap(arr)
		s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:4]
	printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)
	fmt.Println()


	var names2 = []string{"a","b","c","d"}
	names2[0] = "A"
	// names2[4] = "E"  Can't add this way
	fmt.Println("names2 array : ",names2)
	names2 = append(names2, "e","f","g")
	fmt.Println("names2 array : ",names2)
	fmt.Println(names2[2:5]) 

	//dynamic array
	var numbers2 = make([]int,4) //len()=4
	var numbers3 = make([]int,3,5) //len()=3, cap()=5
	numbers2[0] = 1
	numbers2[1] = 2
	numbers2[2] = 3
	numbers2[3] = 4

	numbers3[0]=0
	numbers3[1]=2
	numbers3[2]=3
	fmt.Println("numbers3 array: ",numbers3)
	// numbers3[3]=4
	// fmt.Println("numbers3 array: ",numbers3) RUNTIME ERROR (INDEX OUT OF BOUND)

	

	//Add elements to slice
	numbers2 = append(numbers2,5,6,7,8,9)
	fmt.Println("numbers2 array : ",numbers2)
	// numbers2[9]=10; ERROR

	//Sorting an array
	sort.Ints(numbers2)
	fmt.Println("Sorted numbers2 array : ",numbers2)

	//Remove from slice
	numbers2 = append(numbers2[:2],numbers2[3:]... ) //Remove element at index=2
	fmt.Println("numbers2 array : ",numbers2)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}