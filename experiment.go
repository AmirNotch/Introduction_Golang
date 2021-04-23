package main

import (
	"fmt"
	"time"
)

func main() {
	string1 := "cat"
	// Length is 13 times 5 which is 65.
	string2 := "cat0123456789" +
		"cat0123456789" +
		"cat0123456789" +
		"cat0123456789" +
		"cat0123456789"


	t0 := time.Now()

	for i := 0; i < 100000000; i++ {
		result := len(string1)
		if result != 3 {
			return
		}
	}

	t1 := time.Now()

	for i := 0; i < 10000000000; i++ {
		result := len(string2)
		if result != 65 {
			return
		}
	}
	t2 := time.Now()

	fmt.Println(t1.Sub(t0))
	fmt.Println(t2.Sub(t1))

}
