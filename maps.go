package main
import "fmt"

func main(){
	fmt.Println("Maps")
	languages := make(map[string]string)
	languages["JS"] = "JavaScript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"
	fmt.Println(languages)
	fmt.Println(languages["PY"])

	//delete
	delete(languages,"RB")
	fmt.Println(languages)

	//loops in map
	for key,value := range(languages){
		fmt.Printf("%v : %v\n",key,value)
	}
}