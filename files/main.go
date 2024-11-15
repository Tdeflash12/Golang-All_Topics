package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to the files in golang")
	content:= "Abhesh Mandal us cdds"
	file, err :=os.Create("./myAbhesh.txt");
	// if err != nil{
	// 	panic(err)

	// }

	
	checkNilErr(err)
	length , err :=io.WriteString(file,content)
	checkNilErr(err)
	fmt.Println("Length is:",length)
	defer file.Close()
	readFile("./myAbhesh.txt")
}
func readFile(filname string){
 dataBytype, err :=ioutil.ReadFile(filname)
 checkNilErr(err)
 fmt.Println("Text data in file is \n", string(dataBytype))
}
func checkNilErr(err error)  {
	if err !=nil{
		panic (err)
	}
	
}
