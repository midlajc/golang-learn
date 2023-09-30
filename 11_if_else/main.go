package main

import (
	"fmt"
)

func main() {
	fmt.Println("If Else In Go Lang")

	loginCount := 3
	var message string
	if loginCount == 0 {
		message = "Zero login attempts"
	} else if loginCount == 1 {
		message = "One login attempt"
	} else if loginCount == 2 {
		message = "Two login attempts"
	} else if loginCount == 3 {
		message = "Three login attempts"
	} else {
		message = "Some one tried to login"
	}

	fmt.Println(message)

	if 9%2 == 0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is odd")
	}

	if num := 10; num%2 == 0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is odd")
	}

	// if err != nil {
	// 	fmt.Println(err)
	// }
}
