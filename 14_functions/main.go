package main

import "fmt"

func main() {
	fmt.Println("Functions In Go lang")
	greeter()
	sum := addByValue(10, 20)
	fmt.Println(sum)
	addByReference(2, 5, &sum)
	fmt.Println(sum)
	sum, message := proAdder(1, 2, 3, 4, 5, 6)
	fmt.Println(message, sum)
}

func greeter() {
	fmt.Println("Greetings")
}

func addByValue(num1 int, num2 int) int {
	return num1 + num2
}

func addByReference(num1 int, num2 int, sum *int) {
	*sum = num1 + num2
}

func proAdder(values ...int) (int, string) {
	total := 0
	for _, value := range values {
		total += value
	}
	return total, "hei"
}
