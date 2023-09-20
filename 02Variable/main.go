package main

import "fmt"

// this will not work
// token := "MIDLAJ C"
// this will work
var token = "MIDLAJ C"

// privet variable
const isCool = true

// public variable
const IsCool = true

func main() {
	var username string = "MIDLAJ C"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T\n", username)

	fmt.Println()
	var isLock bool = true
	fmt.Println(isLock)
	fmt.Printf("Variable is of type: %T\n", isLock)

	fmt.Println()
	var smallValue uint8 = 255
	fmt.Println(smallValue)
	fmt.Printf("Variable is of type: %T\n", smallValue)

	fmt.Println()
	var smallFloat float32 = 255.455
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T\n", smallFloat)

	// default values and some aliases
	fmt.Println()
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T\n", anotherVariable)

	fmt.Println()
	var anotherBool string
	fmt.Println(anotherBool)
	fmt.Printf("Variable is of type: %T\n", anotherBool)

	//implicit type
	fmt.Println()
	var implicitVariable = "MIDLAJ C"
	fmt.Println(implicitVariable)
	fmt.Printf("Variable is of type: %T\n", implicitVariable)

	//no var style
	fmt.Println()
	implicitVariable2 := "MIDLAJ C"
	fmt.Println(implicitVariable2)
	fmt.Printf("Variable is of type: %T\n", implicitVariable2)

	//multiple declaration
	fmt.Println()
	var (
		name   = "MIDLAJ C"
		age    = 25
		height int
	)
	fmt.Println(name, age, height)
	height = 5
	fmt.Println(name, age, height)
	fmt.Printf("Variable is of type: %T\n", name)
	fmt.Printf("Variable is of type: %T\n", age)
	fmt.Printf("Variable is of type: %T\n", height)

	fmt.Println()
	var name1 string
	fmt.Println(name1)
	fmt.Printf("Variable is of type: %T\n", name1)
	name1 = "MIDLAJ C"
	fmt.Println(name1)
	fmt.Printf("Variable is of type: %T\n", name1)

	fmt.Println()
	fmt.Println(isCool)
	fmt.Printf("Variable is of type: %T\n", isCool)
}
