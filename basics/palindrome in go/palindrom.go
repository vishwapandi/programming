package main

import (
	"fmt"
)

func isPalindrome(str string, start int, end int) bool {
	if start >= end {
		return true
	}
	if str[start] != str[end] {
		return false
	}
	return isPalindrome(str, start+1, end-1)
}

func main() {
	var str string

	fmt.Print("Enter a string: ")
	fmt.Scan(&str)

	if isPalindrome(str, 0, len(str)-1) {
		fmt.Println("palindrome")
	} else {
		fmt.Println("not a palindrome")
	}
}
