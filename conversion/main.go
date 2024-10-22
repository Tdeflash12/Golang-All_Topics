package main

import (
	"bufio"
	"strings"

	"fmt"
	"os"
	"strconv"
)
func main()  {

	fmt.Println("Welcome to Our shop")
	fmt.Println("Please rate our shop between 1 and 10")
	reader :=bufio.NewReader(os.Stdin)
	  input, _:= reader.ReadString('\n')
	  fmt.Println("Thanks for rating,",input)

	  numRating, err:= strconv.ParseFloat(strings.TrimSpace(input),64)
	  if err != nil {
	fmt.Println(err)
	  }else{
		fmt.Println("Added 1 to yoyr String: ",numRating + 1)
	  }
}