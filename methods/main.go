package main

import "fmt"

func main() {
	fmt.Println("Methods in Golang")

	//no inheritance in golang; No super or parent

	Abhesh := User{"Abhesh","Mandal","abheshmandal249@gmail.com",true,18}
	 fmt.Println(Abhesh)
	 fmt.Printf("Abhesh details given here :%+v\n",Abhesh)
	 // Output Given
	//  Abhesh details given here :{Name:Abhesh Cast:Mandal Email:abheshmandal249@gmail.com Status:true Age:18}
	fmt.Printf("Name is %v and email is %v\n ",Abhesh.Name,Abhesh.Email)
// Output := Name is Abhesh and email is abheshmandal249@gmail.com
Abhesh.GetStatus()
Abhesh.NewMail()
fmt.Printf("Name is %v and email is %v\n ",Abhesh.Name,Abhesh.Email)

}

type User struct {
	Name   string
	Cast   string
	Email  string
	Status bool
	Age    int
}
func(u User) GetStatus(){
	fmt.Println("Is user is active",u.Status)
}
func (u User) NewMail(){
	u.Email= "abheshmandal249@gmail.com"
	fmt.Println("Email of the user is ", u.Email)
}

