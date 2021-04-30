package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"unicode/utf8"
)

/*func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}*/

func pow(x,y,lim float64) float64{
	if v := math.Pow(x, y); v < lim{
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func main(){

	in := bufio.NewScanner(os.Stdin)
	alreadySeen := make(map[string]bool)
	for in.Scan(){
		txt := in.Text()

		if _, found := alreadySeen[txt]; found {
			continue
		}
		alreadySeen[txt] = true

		fmt.Println(txt)
	}

	/*sum := 1
	for ; sum < 1000; {
		sum += sum
	}
	fmt.Println(sum)*/

	/*fmt.Println(sqrt(2),sqrt(-4))*/

	/*fmt.Println(
		pow(3,2,10),
		pow(3,3,20))

	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
*/
	var num0 int

	var num1 int = 1

	var num2 = 20
    fmt.Println(num0, num1, num2)

	num := 30

	fmt.Println(num)

	num += 1
	fmt.Println("+=", num)

	num++
	fmt.Println("++", num)

	userIndex := 10  // Правильно

	user_index := 10  // НЕ Правильно

	fmt.Println(userIndex,user_index)

	var weight, height int = 10, 20

	fmt.Println(weight,height)

	weight, height = 11, 24

	fmt.Println(weight,height)

	weight, myname := 17, 90

	fmt.Println(weight, myname)

	fmt.Println(weight, height,myname)

	var bigInt int64 = 1<<32 -1

	fmt.Println(bigInt)

	var unsignedInt uint = 100500

	fmt.Println(unsignedInt)

	var unsignedBigInt uint64 = 1<<64 - 1

	fmt.Println(unsignedBigInt)

	var pi float32 = 3.141
	var e = 2.718
	goldenRation := 1.618

	fmt.Println(pi,e,goldenRation)

	var b bool
	var isOk bool = true
	var success = true
	cond := true

	fmt.Println(b,isOk,success,cond)

	var c complex128 = -1.1 + 7.12i
	c2 := -1.1 + 7.12i

	fmt.Println(c,c2)

	var str string

	fmt.Println(str)

	var hello string = "Hello\n\t"

	fmt.Println(hello)

	var world string = "World\n\t"

	fmt.Println(world,hello)

	var helloWorld = "Hello, World!"

	var rawBinary byte = '\x27'

	fmt.Println(helloWorld,rawBinary)

	helloworld := "Привет Мир"

	andGoodMorning := helloworld + " и доброе утро!"

	fmt.Println(andGoodMorning)

	bytelen := len(helloworld)
	symbols := utf8.RuneCountInString(helloworld+"+123")

	fmt.Println(bytelen,symbols)


	Hello := helloworld[:19]
	H := helloworld[0]

	fmt.Println(Hello,H)

	var byteString = []byte(helloworld)
	helloWorld = string(byteString)

	var crypt = string(H)
	fmt.Println(byteString,helloworld,crypt)

	fmt.Println(hello1)

	var hellow2orld string = "hello"
	fmt.Println(hellow2orld)
	fmt.Println(len(hellow2orld))
	fmt.Println(utf8.RuneCountInString(hellow2orld))

	var I int
	for i := 0;i <= len(hellow2orld); i++{
		I = i
	}
	fmt.Println(I)

	byteelen := "foobar"
	fmt.Println(len(byteelen))

	const (
		one = iota + 1
		two
		three
	)
	fmt.Println(one,two,three)


	var buf []int
	buf = append(buf, 2)
	buf = append(buf,15)
	fmt.Println(buf)

	otherBuf := make([]int,3)
	otherBuff := make([]int,3)
	otherBuf = append(otherBuf, 4)
	fmt.Println(otherBuf)
	buf = append(buf, otherBuf...)
	buf = append(buf, otherBuff...)
	fmt.Println(cap(buf))

}

const pi = 3.141
const (
	hello1 = "Hello21"
	e = 2.817
)
const (
	zero = iota
	_
	three
)
