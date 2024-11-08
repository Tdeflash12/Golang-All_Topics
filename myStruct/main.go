package main

import "fmt"

func main() {
	fmt.Println("Structs in Golang")

	//no inheritance in golang; No super or parent

	Abhesh := User{"Abhesh","Mandal","abheshmandal249@gmail.com",true,18}
	 fmt.Println(Abhesh)
	 fmt.Printf("Abhesh details given here :%+v\n",Abhesh)
	 // Output Given
	//  Abhesh details given here :{Name:Abhesh Cast:Mandal Email:abheshmandal249@gmail.com Status:true Age:18}
	fmt.Printf("Name is %v and email is %v ",Abhesh.Name,Abhesh.Email)
// Output := Name is Abhesh and email is abheshmandal249@gmail.com
}

type User struct {
	Name   string
	Cast   string
	Email  string
	Status bool
	Age    int
}
