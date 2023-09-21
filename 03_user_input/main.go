package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Enter a number: ")
	// using Scanf form package fmt to get user input
	var input float64
	fmt.Scanf("%f", &input)

	output := input * 2

	fmt.Println(output)

	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Enter second number: ")
	input2, _ := reader.ReadString('\n')

	fmt.Printf("Input %s", input2)

}
