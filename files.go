package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewScanner(os.Stdin)
	alreadySeen := make(map[string]bool)
	for in.Scan() {
		txt := in.Text()

		if _, found := alreadySeen[txt]; found {
			continue
		}

		alreadySeen[txt] = true

		fmt.Println(txt)
	}
}

/*package main
import (
	"fmt"
	"os"
	"io"
)

func main() {
	file, err := os.Open("data_map.txt")
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	data := make([]byte, 64)

	alreadySeen := make(map[string]bool)

	for{

		_, txt := file.Read(data)

		if _, found := alreadySeen[txt]; found {
			continue
		}

		alreadySeen[txt] = true

		fmt.Println(txt)
	}
}*/