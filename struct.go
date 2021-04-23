package main

import "fmt"

type Account struct {
	Id       int
	Name     string
	Cleaner  func(string) string
	Owner    Person
	Person
}
type Person struct {
	Id      int
	Name    string
	Address string
}
func main() {
	var acc Account = Account{
		Id:  1,
		Name: "Burguy",
		Person: Person{
			Name: "Georgiy",
			Address: "Hollywood",
		},
	}
	fmt.Printf("%#v\n\n\n", acc)

	acc.Owner = Person{2, "Kudbidin", "Negmat Karabaeva"}
	fmt.Printf("%#v\n\n", acc)

	fmt.Printf("%#v\n",acc)


	fmt.Println(acc.Name)
	fmt.Println(acc.Person.Name)
}
