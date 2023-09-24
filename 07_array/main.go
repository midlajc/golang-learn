package main

import "fmt"

func main() {
	fmt.Println("Arrays in Go")

	myStringArray := [5]string{"a", "b", "c"}
	myStringArray[4] = "d"

	fmt.Println("Value of myStringArray is", myStringArray)
	fmt.Println("Length of myStringArray is", len(myStringArray))

	var myArray [4]int = [4]int{1, 2, 3}

	fmt.Println("Value of myArray is", myArray)
	fmt.Println("Length of myArray is", len(myArray))
	fmt.Println("Capacity of myArray is", cap(myArray))
	fmt.Println("Type of myArray is", fmt.Sprintf("%T", myArray))
}
