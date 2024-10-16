package main

import "fmt"

func getFullName() (string, string) {
	return "kocak", "gaming"
}

func main() {
	//firstName, lastName := getFullName()
	//fmt.Println(lastName, firstName)

	firstName, _ := getFullName()
	fmt.Println(firstName)
}
