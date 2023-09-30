package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch in Go Lang")

	rand.Seed(time.Now().UnixNano())

	diceNumber := rand.Intn(6) + 1

	switch diceNumber {
	case 1:
		fmt.Println("Game Started, Dice number is 1")
	case 2:
		fmt.Println("Dice number is 2")
	case 3:
		fmt.Println("Dice number is 3")
	case 4:
		fmt.Println("Dice number is 4")
	case 5:
		fmt.Println("Dice number is 5")
	default:
		fmt.Println("Dice number is 6")
	}
}
