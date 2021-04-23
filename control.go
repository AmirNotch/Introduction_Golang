package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	boolVal := true
	if boolVal {
		fmt.Println("Hello World!")
	}

	mapVal := map[string]string{"name": "Amir"}

	if keyValue, keyExist := mapVal["name"]; keyExist {
		fmt.Println("name", keyValue)
	}

	if _, keyExist := mapVal["name"]; keyExist {
		fmt.Println("key 'name' exist")
	}

	cond := 1

	if cond == 1 {
		fmt.Println(1)
	} else if cond == 2 {
		fmt.Println(2)
	}

	strVal := "name"
	switch strVal {
	case "name":
		fallthrough
	case "test", "lastName":
		fmt.Println("hi Nigga")
	default:

	}

loop:
	for key, val := range mapVal {
		println("switch in loop", key, val)
		switch {
		case key == "lastName":
			break
			println("dont show")
		case key == "name" && val == "Amir":
			println("hi find me")
			break loop
			println("good buy")

		}
	}

	str := "Hello world"
	for pos, char := range str {
		fmt.Printf("%#U at pos %d\n\n", char, pos)
	}

	/*var acc Account = Account{
		Id:   1,
		Name: "Nigga",
	}
	fmt.Printf("Hello\n", acc)*/

	/*acc.Owner = Person{2, "Nizor", "Dushanbe"}*/

/*	fmt.Printf("Hello\n", acc)*/

	sl := []int{1, 3, 4}
	idx := 0

	for idx < len(sl) {
		fmt.Println("ind", idx, "value", sl[idx])
		idx++
	}

	for ind := range sl {
		fmt.Println("ind", ind, "value", sl[ind])
	}

	for ind, val := range sl {
		fmt.Println("ind", ind, "value", val)
	}

	for i := 0; i < len(sl); i++ {
		fmt.Println("ind", i, "value", sl[i])
	}
	for hin := range sl {
		fmt.Println("\n", "honey", sl[hin])
	}

	straniza := "Ky MNR"
	for pos, val := range straniza {
		fmt.Println(pos, utf8.RuneCountInString(string(val)))
	}

	fmt.Println()

	profile := map[int]string{1: "Vasya", 2: "Kudrat"}
	for key := range profile {
		fmt.Println(key)
	}
	for key, val := range profile {
		fmt.Println(key, val)
	}
	for _, val := range profile {
		fmt.Println(val)
	}

	fmt.Println(singleIn(2))
	fmt.Println(multipleIn(2, 2, 2))
	fmt.Println(result())
	fmt.Println(multipleResult(3))
	fmt.Println(fact(5))
	fmt.Println(multipleNamedReturn(true))
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(nums, sum(nums...))

	func(in string) {
		fmt.Println("Hello", in)
	}("Kdbidin")

	printer := func(in string) {
		fmt.Println("Printer outs:", in)
	}
	type strFuncType func(string)

	worker := func(callback strFuncType) {
		callback("as callback")
	}
	worker(printer)

	prefixer := func(prefix string) strFuncType {
		return func(in string) {
			fmt.Printf("[%s] %s\n\n", prefix, in)
		}
	}

	successLogger := prefixer("SUCCESS")
	successLogger("expected behavior")
}

func singleIn(in int) int {
	return in
}

func multipleIn(a,b int,c int) int {
	return a + b + c
}
func result() (out string) {
	out = "helicopter"
	return
}

func multipleResult(in int) (int, error){
	if in > 2 {
		return 0, fmt.Errorf("bycicle")
	}
	return 0, nil
}

func fact(x int) int {
	if x == 0 {
		return 1
	}
	return x * fact(x-1)
}

func multipleNamedReturn(ok bool) (rez int, err error){
	if ok {
		err = fmt.Errorf("Some error happend")
		return
	}
	rez = 2
	return
}

func sum(in ...int) (result int) {
	fmt.Println("in :=", in)
	for _, val := range in {
		result += val
	}
	return
}

/*type Account struct {
	Id       int
	Name     string
	LastName string
	Cleaner func(string)string
	Owner    Person
	Person

}*/
/*type  Person struct {
	Id      int
	Name    string
	Address string
}*/

