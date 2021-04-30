package main

import (
	"fmt"
	"os"
)

func main() {
	text := "Hello Bold"
	file, err := os.Create("hello.txt")

	if err != nil {
		fmt.Println("Unable to create files:", err)
		os.Exit(1)
	}
	defer file.Close()
	file.WriteString(text)

	fmt.Println("Done")
}
