package main

import "fmt"

type UserID int

type UserName string

func main() {

	const (
		year = 2017

		yearTyped int = 2017
	)

	var month int32 = 13
	fmt.Println(month + year)


    idx := 2
    var uid UserID = 42

    myID := UserID(idx)

    fmt.Println(uid,myID)

    a := 2
    b := &a
    *b = 3
    c:= &a

fmt.Println(b)
    d := new(int)
    *d = 12
    *c = *d
    *d = 13
fmt.Println(c)
    c = d
    *c = 14
    fmt.Println(c)

    idx3 := "HIII"
    var Num UserName = "Hellli"
    Num2 := UserName(idx3)
    println(Num, Num2)
    fmt.Println(Num, Num2)
}
