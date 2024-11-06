package main

import "fmt"

func main() {
	fmt.Println("Welcome to the Array in Golang")
	var CarList [5] string
	CarList [0] = "Bugatii"
	CarList [1] = "Marcities"
	CarList [2] = "Farriri"
	CarList [3] = "Lamborgini"
	CarList [4] = "BMW"
	


fmt.Println("The list of Cars is",CarList)
// lens is used to calculate overall all number of products that is declared in an Array
fmt.Println("The list of Cars is",len(CarList))
// Be careful about giving index number 
var vegitableList = [5] string {"potato", "cabbage"}
fmt.Println("Vegetable list is ",vegitableList)
fmt.Println("Vegetable list is ",len(vegitableList) )

}