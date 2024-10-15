package main

import "fmt"

func main() {
	var name string
	name = "Maulana Givari"
	fmt.Println(name)

	name = "Muhammad Maulana"
	fmt.Println(name)

	var (
		firstName  = "Muhammad"
		middleName = "Maulana"
		lastName   = "Givari"
	)

	fmt.Println(firstName)
	fmt.Println(middleName)
	fmt.Println(lastName)
	fmt.Println(firstName + " " + middleName + " " + lastName)

}
