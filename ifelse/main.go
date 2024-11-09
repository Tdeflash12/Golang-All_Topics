package main

import "fmt"

func main() {
	fmt.Println("IfElse in Golang")
	loginCount := 10
	var result string
	if loginCount < 10 {
		result = "Regular user Done "
	} else if loginCount > 10 {
		result = "Done"
	} else {
		result = "Exactly 10 login count"
	}
	fmt.Println(result)
	// To check number is even or odd
	if 10%2 == 0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is odd")
	}
	// To check Number is less than 10 Or not BY Abhesh Mandal
	if num := 10; num <= 10 {
		fmt.Println("Number is less than 10")
	} else {
		fmt.Println("Numbe is more than 10")
	}

}
