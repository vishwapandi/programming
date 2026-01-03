package main

import "fmt"

func leap_year(year int) bool {
	if year%400 == 0 {
		return true
	} else if year%100 == 0 {
		return false
	} else if year%4 == 0 {
		return true
	}
	return false
}

func main() {
	var year int

	fmt.Print("Enter a year: ")
	fmt.Scan(&year)

	if leap_year(year) {
		fmt.Println("Leap Year")
	} else {
		fmt.Println("Not a Leap Year")
	}
}
