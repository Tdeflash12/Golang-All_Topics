package main

import "fmt"

func main() {
	
	op()
	fmt.Println("Welcome to the function in golang")
	result := sum(3,9)
	fmt.Println("The result is ", result)
	proResult, message:= proAddder(2,34,43,2)
	fmt.Println("pro result is ",proResult)
	fmt.Println("oh vai",message)
}

func op()  {
	fmt.Println("Namastey From Abhesh Mandal")
}
func sum(valueOne int, valueTwo int) int {
	return valueOne + valueTwo
	
}


func proAddder(values ...int) (int,string){
total:= 0
for _, val:= range values{
	total += val
}
return total,"Abhesh Mandal"
}
