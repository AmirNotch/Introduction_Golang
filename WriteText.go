package main

import (
	"fmt"
	"os"
)

func main() {
	data := []byte("hello Bold")
	file, err := os.Create("hello.txt")
	if err != nil {
		fmt.Println("Unable to create files:", err)
		os.Exit(1)
	}
	defer file.Close()
	file.Write(data)

	fmt.Println("Done")
}
