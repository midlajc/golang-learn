package main

import "fmt"

func main() {
	fmt.Println("Loops in Go Lang")

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	}

	for i := range days {
		fmt.Println(days[i])
	}

	for index, day := range days {
		fmt.Printf("Index: %v, Day: %v\n", index, day)
	}

	for _, day := range days {
		fmt.Printf("Day: %v\n", day)
	}

	for i := 1; i <= 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Printf("%v ", i)
		if i == 8 {
			fmt.Println()
			break
		}
		if i == 10 {
			fmt.Println()
		}
	}

	i := 0
lab:
	fmt.Println("Inside Label ", i)

	if i < 10 {
		i++
		goto lab
	}
}
