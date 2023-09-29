package main

import (
	"fmt"
)

func main() {
	fmt.Println("Maps in Go")

	languages := make(map[string]string)
	fmt.Println(languages)
	languages["js"] = "Javascript"
	languages["rb"] = "Ruby"
	languages["py"] = "Python"
	languages["go"] = "Golang"
	languages["ts"] = "Typescript"
	languages["cs"] = "C#"

	fmt.Println(languages)
	fmt.Println(languages["rb"])
	fmt.Printf("Languages: %T\n", languages)
	fmt.Println("Languages Length: ", len(languages))

	// delete from map
	delete(languages, "cs")
	fmt.Println(languages)

	// check if key exists
	value, ok := languages["py"]
	fmt.Println(value, ok)
	value, ok = languages["test"]
	fmt.Println(value, ok)

	// iterate over map
	for key, value := range languages {
		fmt.Printf("Key: %v, Value: %v\n", key, value)
	}
}
