package main

import "fmt"

func main() {
	fmt.Println("Slices in Go")

	myStringSlice := []string{"a", "b", "c"}
	fmt.Println("Value of myStringSlice is", myStringSlice)
	fmt.Println("Length of myStringSlice is", len(myStringSlice))
	myStringSlice = append(myStringSlice, "d")
	fmt.Println("Length of myStringSlice is", len(myStringSlice))
	fmt.Println("Capacity of myStringSlice is", cap(myStringSlice))
	myStringSlice = append(myStringSlice, "d")
	fmt.Println("Length of myStringSlice is", len(myStringSlice))
	fmt.Println("Capacity of myStringSlice is", cap(myStringSlice))
	myStringSlice = append(myStringSlice, "d")
	fmt.Println("Length of myStringSlice is", len(myStringSlice))
	fmt.Println("Capacity of myStringSlice is", cap(myStringSlice))
	myStringSlice = append(myStringSlice, "d")
	fmt.Println("Length of myStringSlice is", len(myStringSlice))
	fmt.Println("Capacity of myStringSlice is", cap(myStringSlice))
	fmt.Println("Type of myStringSlice is", fmt.Sprintf("%T", myStringSlice))
}
