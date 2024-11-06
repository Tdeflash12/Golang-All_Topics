package main

import (
	"fmt"
	"sort"
)

func main()  {
	fmt.Println("Welcome to the new Topics of Slices")
	var CarList = []string{"BMW","Bugatti","Lamborgini"}
	fmt.Printf("Types of car list is %T\n",CarList)
	CarList = append(CarList,"Suzuki","Rollce Roayles")
	fmt.Println(CarList)
	// append is used to add products or valuess
	CarList= append(CarList[:3])
	fmt.Println(CarList)

	highScore:= make([]int,4)
	highScore[0] =234
	highScore[1] =999
	highScore[2] =500
	highScore[3] =900
	//highScore[4] =777
	highScore =append(highScore, 555,666,777)
	fmt.Println(highScore)
	// sort.Ints is used to make arrangement in increasing order
	sort.Ints(highScore)
	fmt.Println(highScore)
// sort.IntsAreSorted are used to check whether the value is in increasng order or not
	fmt.Println(sort.IntsAreSorted(highScore))



}