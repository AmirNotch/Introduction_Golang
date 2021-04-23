package main

import "fmt"

func main() {
	/*defer fmt.Println("After Work")
	defer fmt.Println(getSomeResult())
	defer fmt.Println(getSomeResult())
	defer fmt.Println(getSomeResult())
	fmt.Println("Compiler is working")*/
	deferTest()
}

func deferTest(){
	defer func(){
		if err := recover(); err != nil {
			fmt.Println("panic happened FIRST", err)
		}
	}()
	defer func(){
		if err := recover(); err != nil {
			fmt.Println("panic happened SECOND", err)
			panic(2)
		}
	}()
	fmt.Println("Some useful work")
	panic("something happened")
	return
}


/*func getSomeResult() string {
	fmt.Println(2+2*2*2)
	return "I'll be back"
}*/
