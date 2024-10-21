package main

import (
	"bufio"
	"fmt"
	"os"
)
func  main ()  {
	welcome := "wELCOME TO THE ABHESH MANDAL USER INPUT"
	fmt.Println(welcome)
	

	reader:= bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating of our pizaa:")

	// comma ok // err ok

	input, _:= reader.ReadString('\n');
	fmt.Println("Thanks for rating, ",input)
	fmt.Printf("Types of this rating is %T ",input)

}
