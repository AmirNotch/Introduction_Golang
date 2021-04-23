package main

import "fmt"

func main() {
	buf := []int{1,2,3,4,5}
	fmt.Println(buf)

	sl1 := buf[1:4]
	sl2 := buf[:2]
	sl3 := buf[2:]
	fmt.Println(sl1,sl2,sl3)

	newBuf := buf[:]
	fmt.Println(newBuf)

	newBuf[0] = 9
	fmt.Println(newBuf)
	fmt.Println(buf)

	newBuf = append(buf, 9)

	fmt.Println(newBuf)

	newBuf = make([]int, len(buf),cap(buf))
	copy(newBuf,buf)

	fmt.Println(newBuf)

	ints := []int{1,2,3,4}
	copy(ints[1:3], []int{5,6})
	fmt.Println(ints)

	slice := make([]int,12,12)
	fmt.Println(slice)
	for i := 0; i < len(slice); i++{
		slice[i] = i
	}
	fmt.Println(slice)
	i1:=0
	for i := 1; i == 10; i++{
		i1++
	}
	slice2 := make([]int, i1, i1)
	fmt.Println(slice2)
	for i := 0; i < 10; i++{
		slice2[i] = i
	}
	fmt.Println(slice2)
}
