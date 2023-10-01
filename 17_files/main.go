package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Files Handling in Go Lang")
	content := "Hello World!"

	file, err := os.Create("./test.txt")
	// if err != nil {
	// 	panic(err)
	// }
	checkNilError(err)

	file.WriteString(content)
	file.WriteString("\n")
	length, err := file.WriteString("New Line")
	checkNilError(err)

	fmt.Println("Length is: ", length)
	defer file.Close()
	readFile("./test.txt")
}

func readFile(fileName string) {
	// file, err := os.Open(fileName)
	// checkNilError(err)
	// defer file.Close()
	// fileInfo, err := file.Stat()
	// checkNilError(err)
	// fileSize := fileInfo.Size()
	// buffer := make([]byte, fileSize)
	// _, err = file.Read(buffer)
	// checkNilError(err)
	// content := string(buffer)
	// fmt.Println(content)

	dataByte, err := ioutil.ReadFile(fileName)
	checkNilError(err)

	fmt.Println("File Content Is:\n", string(dataByte))

}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
