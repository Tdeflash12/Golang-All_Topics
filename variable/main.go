package main

import "fmt"
const LoginToken string ="Mandal"

func main() {
	
	var username string = "Abhesh Mandal"
	fmt.Println(username)
	fmt.Printf("variable is type : %T \n", username)

	var isAbhesh bool = false
	fmt.Println(isAbhesh)
	fmt.Printf("variable is type : %T \n",isAbhesh)

	var smallValue uint8 = 250
	fmt.Println(smallValue)
	fmt.Printf("variable is type : %T \n",smallValue)

	var smallfloat float64 = 250.8565668786
	fmt.Println(smallfloat)
	fmt.Printf("variable is type : %T \n",smallfloat)

	//default values and some aliases
	var anotheerVariable int
	fmt.Println(anotheerVariable)
	fmt.Printf("Variable is of type :%T \n",anotheerVariable)

	// implict types
	var name ="Abhesh"
	fmt.Println(name)

	// no var style
	numberOfUser:=30000
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type :%T \n",LoginToken)

}
