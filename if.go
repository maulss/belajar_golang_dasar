package main

import "fmt"

func main() {
	name := "maulana givari"
	if name == "maulana" {
		fmt.Println("hello", name)
	} else if name == "maulana givari" {
		fmt.Println("hello if else ", name)
	} else {
		fmt.Println("salah nama", name)
	}

	if length := len(name); length < 3 {
		fmt.Println("length is to shrot", length)
	} else {
		fmt.Println("length is too long", length)
	}
}
