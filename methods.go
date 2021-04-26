package main

import "fmt"

type Person1 struct {
	Id  int
	Name string
}

func (p Person1) UpdateName(name string){
	p.Name = name
}

func (p *Person1) SetName(name string) {
	p.Name = name
}

type Account1 struct {
	Id int
	Name string
	Person1
}

type mySlice []int

func (sl *mySlice) Add(val int){
	*sl = append(*sl, val)
}

func (sl *mySlice) Count() int{
	return len(*sl)
}

func main() {
	/*pers := Person1{1, "Nikita"}*/
	pers := new(Person1)
	fmt.Printf("%#v\n",pers)
	pers.SetName("Hello Debili")
	//fmt.Printf("%#v",pers)
	//(&pers).SetName("Hollywood")
	fmt.Printf("%#v",pers)

	var acc Account1 = Account1{
		Id: 1,
		Name: "Rubinsky",
		Person1: Person1{
			Id: 2,
			Name: "Drimsky",
		},
	}

	acc.Person1.SetName("Ibragimov Amir")

	fmt.Printf("%#v \n", acc)

	sl := mySlice([]int{1,2,3})
	sl.Add(3)
	fmt.Println(sl.Count(), sl)
}
