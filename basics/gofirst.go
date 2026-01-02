package main

import "fmt"

func main(){
	var num int
	fmt.printf("Enter a number :")
	fmt.scanf("%d",&num)
	if num % 2 == 0 {
		fmt.printf("%d is Even", num)
	} else {
		fmt.printf("%d is odd" , num)
	}
}