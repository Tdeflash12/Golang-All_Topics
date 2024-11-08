package main

import "fmt"

func main() {
	fmt.Println("Welcome To the Topics Of Pointers")

	// var ptr *int 
	// // fmt.Println("Value of pointet is :",ptr)
	myNumber:= 23
	var ptr = &myNumber
	fmt.Println("Value of the actual pointer is :",ptr)
	fmt.Println("Value of the actual pointer is :",*ptr)
	*ptr = *ptr * 2
	fmt.Println("Value of new pointer is",myNumber)

}
