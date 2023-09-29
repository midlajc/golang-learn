package main

import (
	"fmt"
	"sort"
)

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

	fruitList := []string{"Apple", "Tomato", "Peach"}
	fruitList = append(fruitList, "Mango", "", "Banana")

	fmt.Println("Fruit List:", fruitList)

	fruitList = append(fruitList[1 : len(fruitList)-1])
	// fruitList = append(fruitList[1:3])
	// fruitList = append(fruitList[:3])

	fmt.Println("Fruit List:", fruitList)

	highScores := make([]int, 4)

	fmt.Println("High Scores:", highScores)
	fmt.Println("Length of High Scores:", len(highScores))

	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 465
	highScores[3] = 867
	fmt.Println("High Scores:", highScores)
	// highScores[4] = 777 // This will throw an error

	highScores = append(highScores, 555, 666)

	fmt.Println("High Scores:", highScores)
	fmt.Println("Length of High Scores:", len(highScores))
	fmt.Println("Capacity of High Scores:", cap(highScores))

	highScores = []int{234, 945, 465, 867, 555, 666}
	// highScores[7] = 999 // This will throw an error

	fmt.Println("High Scores:", highScores)
	fmt.Println("Length of High Scores:", len(highScores))
	fmt.Println("Capacity of High Scores:", cap(highScores))

	highScores = append(highScores, 777)
	fmt.Println("High Scores:", highScores)
	fmt.Println("Length of High Scores:", len(highScores))
	fmt.Println("Capacity of High Scores:", cap(highScores))
	fmt.Println("Is High Scores sorted?", sort.IntsAreSorted(highScores))

	sort.Ints(highScores)
	fmt.Println("High Scores:", highScores)
	fmt.Println("Length of High Scores:", len(highScores))
	fmt.Println("Capacity of High Scores:", cap(highScores))
	fmt.Println("Is High Scores sorted?", sort.IntsAreSorted(highScores))

	courses := make([]string, 5)
	courses[0] = "Docker"
	courses[1] = "Puppet"
	courses[2] = "Python"
	courses[3] = "Ansible"
	courses[4] = "React"
	fmt.Println("Courses:", courses)
	fmt.Println("Length of Courses:", len(courses))
	fmt.Println("Capacity of Courses:", cap(courses))

	removeIndex := 2

	// remove item without preserving order
	courses = append(courses[:removeIndex], courses[removeIndex+1:]...)
	fmt.Println("Courses:", courses)
}
