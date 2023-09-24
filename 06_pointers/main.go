package main

import "fmt"

func main() {
	fmt.Println("Pointers in Go")

	myNumber := 23

	ptr := &myNumber

	ptr2 := &ptr

	fmt.Println("Value of myNumber is", myNumber)
	fmt.Println("Address of myNumber is", &myNumber)
	fmt.Println("Value of ptr is", ptr)
	fmt.Println("Value of *ptr is", *ptr)
	fmt.Println("Address of ptr is", &ptr)
	fmt.Println("Value of ptr2 is", ptr2)
	fmt.Println("Value of *ptr2 is", *ptr2)
	**ptr2++
	fmt.Println("Value of **ptr2 is", **ptr2)
	fmt.Println("Value of myNumber is", myNumber)

	fmt.Println("Type ptr is", fmt.Sprintf("%T", ptr))
	fmt.Println("Type *ptr is", fmt.Sprintf("%T", *ptr))
	fmt.Println("Type ptr2 is", fmt.Sprintf("%T", ptr2))
	fmt.Println("Type &myNumber is", fmt.Sprintf("%T", &myNumber))
}
