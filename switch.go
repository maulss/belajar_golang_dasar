package main

import "fmt"

func main() {
	name := "rawr"

	switch name {
	case "jaki":

		fmt.Println("name jaki")
	case "maulana":
		fmt.Println("name maulana")
	default:
		fmt.Println("name default")

	}

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("length true")
	case false:
		fmt.Println("length false")
	}

}
