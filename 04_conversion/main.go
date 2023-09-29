package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our Pizza App")
	fmt.Printf("Please rate our pizza between 1 and 5: ")

	reader := bufio.NewReader(os.Stdin)

	rating, _ := reader.ReadString('\n')

	fmt.Println("Thank you for rating our pizza, ", rating)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(rating), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to your rating : ", numRating+1)
	}

	const num = 3
	numStr := strconv.Itoa(num)
	fmt.Printf("numStr, %T\n", numStr)
}
