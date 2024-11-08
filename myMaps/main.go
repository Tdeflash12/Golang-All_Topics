package main

import "fmt"

func main() {
	fmt.Println("Maps in Golang!.")
	languages:= make(map[string]string)
	languages["rs"]="Rust"
	languages["js"]="JavaScript"
	languages["py"]="Python"

	fmt.Println("list of all programming lanugaes",languages)
	fmt.Println("py FullForm is : :",languages["py"])
// delete keyword is used tp delete some values .
	delete(languages,"rs")
	fmt.Println("list of all programming lanugaes",languages)

	//loops are more interesting in golang
	for key,value:=range languages{
			 fmt.Printf("For key %v,value is %v\n",key ,value)
	}
}