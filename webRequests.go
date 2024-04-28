package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url string = "https://somanistoneminerals.com"

func main(){
	response, err := http.Get(url)
	checkError(err)
	fmt.Printf("Response dtype : %T",response)
	fmt.Println("Response : %v",response)

	defer response.Body.Close() //IMP ALways close the connection after response
	dataBytes,err := ioutil.ReadAll(response.Body)
	checkError(err)
	fmt.Println("Response : ",string(dataBytes))
}

func checkError(err error){
	if(err!=nil){
		panic(err)
	}
}