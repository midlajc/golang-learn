package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to Time Study")

	presetTime := time.Now()
	fmt.Println(presetTime)
	fmt.Println(presetTime.Format(time.RFC3339))
	fmt.Println(presetTime.Format("01-02-2006 15:04:05 Monday"))

	fmt.Println()

	createTime := time.Date(2002, time.January, 16, 12, 0, 0, 0, time.Local)
	fmt.Println(createTime)
	fmt.Println(createTime.Format(time.RFC3339))
	fmt.Println(createTime.Format("01-02-2006 15:04:05 Monday"))
}
