package main

import (
	"fmt"
)

func main() {

	var a4 [3]int
	fmt.Println("a1 short", a4)
	fmt.Println("a1 short %v\n", a4)
	fmt.Println("a1 full %#v\n", a4)

	const size = 2
	var a2 [2 * size]bool
	fmt.Println("a2", a2)

	a3 := [...]int{1,2,3,4}
	fmt.Println("a3", a3)

	a3[3] = 23

	fmt.Println(a3)


	a := 2
	b := &a
	*b = 3
	c := &a
	fmt.Println(*c)



	var a1 [5]float64
	a1[0] = 98
	a1[1] = 93
	a1[2] = 77
	a1[3] = 82
	a1[4] = 83

	var total float64 = 0
	fmt.Println(a1[0:])
	for i := 0;i < len(a1); i++ {
		total += a1[i]
	}

	fmt.Println(total / float64(len(a1)))

	total = 0
	for _, value := range a1 {
		total += value
	}
	fmt.Println(total / float64(len(a1)))

	dwa := [...]float64{54,23,5,77,0,4213,45}
	fmt.Println(dwa)

	slice := []int{1,2,3,4}
	slice2 := append(slice,5,6)
	fmt.Println(slice,slice2)

	slice3 := []int{1,2,3}
	slice4 := make([]int,5)
	copy(slice4,slice3)
	fmt.Println(slice3,slice4)


}
