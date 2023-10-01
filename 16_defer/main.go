package main

func main() {
	defer println("in main")
	println("in main 2")
	deferFunction()
}

func deferFunction() {
	for i := 0; i < 5; i++ {
		defer println(i)
	}
}
