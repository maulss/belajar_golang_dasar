package main

import "fmt"

func getHello(name string) string {
	return "hello" + name
}

func main() {
	fmt.Println(getHello("kocak"))
	fmt.Println(getHello("gaming"))
}
