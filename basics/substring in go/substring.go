package main

import (
	"fmt"
	"strings"
)

func isSubstring(str1, str2 string) bool {
	return strings.Contains(str1, str2)
}

func main() {
	var str1, str2 string

	fmt.Print("Enter first String: ")
	fmt.Scan(&str1)

	fmt.Print("Enter Second String: ")
	fmt.Scan(&str2)

	if isSubstring(str1, str2) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
