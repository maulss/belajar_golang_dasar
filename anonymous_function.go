package main

import "fmt"

type Blacklist func(name string) bool

func registerUser(name string, blaclist Blacklist) {
	if blaclist(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

func main() {
	blacklist := func(name string) bool {
		return name == "Anjing"
	}
	registerUser("maulana", blacklist)

	registerUser("Anjing", func(name string) bool {
		return name == "Anjing"
	})
}
