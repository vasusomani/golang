package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main(){
	fileAddress := "./firstGoFile.txt"
	file := createFile(fileAddress)
	writeFile(file,"Hey!! this is my first golang text file")
	fmt.Println(openFile(fileAddress))
}

func createFile(address string) (*os.File){
	file,err := os.Create(address) //File path
	checkError(err)
	fmt.Println("File created successfully!")
	return file
}

func writeFile(file *os.File, content string){
	length,err := io.WriteString(file,content)
	checkError(err)
	fmt.Println("Length of the data is : ",length)
}

func openFile(filePath string)string{
	dataByte,err := ioutil.ReadFile(filePath)
	checkError(err)
	return string(dataByte)
}

func checkError(err error){
	if(err!=nil){
		panic(err) //Stops the execution of program and returns err
	}
}