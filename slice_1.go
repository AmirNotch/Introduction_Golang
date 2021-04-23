package main

import "fmt"

func main() {
	var buf0 []int
	buf1 := []int{}
	buf2 := []int{42}
	buf3 := make([]int,0)
	buf4 := make([]int,6)
	buf5 := make([]int,5,9)

	fmt.Println(buf0,buf1,buf2,buf3,buf4,buf5)

	someInt := buf2[0]

	someInt = 2
	fmt.Println(someInt)

	var buf []int
	buf = append(buf, 9,10)
	buf = append(buf,12)
	fmt.Println(buf)

	otherbuf := make([]int, 3)
	buf = append(buf, otherbuf...)
	buf[4] = 2
	fmt.Println(buf, otherbuf)

	var bufLen, bufCap int = len(buf),cap(buf)

	fmt.Println(bufLen,bufCap)

}
