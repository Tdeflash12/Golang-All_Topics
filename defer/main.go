package main

import "fmt"

func main() {
	// Last in FirstOut
	defer fmt.Println("hello")
	defer fmt.Println("World")
defer fmt.Println("Abhesh Mandal")

	fmt.Println("Welcome to the defer in golang")
	myDefer()
	
}
func myDefer()  {
	for i:=0;i<5;i++{
		defer fmt.Println(i)
	}
}