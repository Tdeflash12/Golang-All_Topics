package main

import "fmt"

func main() {
	fmt.Println("Welocome to the loops in golang");
	days:= [] string{"Sunday","Monday","Tuesday","Wednesday","Thursaday","Friday","Saturday"}
	fmt.Println(days)
	// for i:=0; i<len(days); i++{
	// 	fmt.Println(days[i])
	// }

	// for d :=range days{
	// 	fmt.Println(days[d])
	// }
	// for index,  day:= range days{
	// 	fmt.Printf("index is %v and value is %v\n",index,day)
	// }
	roughValue:= 1

	for roughValue<=10{

		if roughValue ==4{
			goto loop
		}
		if roughValue==5 {
		break
			
		}
		fmt.Println("value is:",roughValue)
		roughValue++;
	}
	loop:
	fmt.Println("Jumping at he Mandal House")
}