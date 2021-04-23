package main

import (
	"fmt"
)

func main() {
	var user map[string]string = map[string]string{
		"name":     "Amir",
		"lastName": "Ibragimov",
	}
	fmt.Println(user)

	profile := make(map[string]string, 10)

	mapLength := len(user)

	fmt.Println(mapLength,profile)

	mName := user["middleName"]
	fmt.Println("mName", mName)

	mName, mNameExist := user["middleName"]
	fmt.Println("mName",mName,"mNameExist",mNameExist)

	_, mNameExist2 := user["middleName"]
	fmt.Println("mNameExist2",mNameExist2)

	delete(user, "lastName")
	fmt.Println(user)

	_, mNameExist3 := user["lastName"]
	fmt.Println("mNameExist3",mNameExist3)


}
