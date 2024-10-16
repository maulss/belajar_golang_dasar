package main

import "fmt"

func goodBye(name string) string {
	return "goog bye" + name
}

func main() {
	googBye := goodBye
	fmt.Println(googBye("goog bye"))
}
