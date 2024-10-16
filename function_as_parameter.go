package main

import "fmt"

type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	fmt.Println("hello", filter(name))
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "...."
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter("Kuy Gaming", spamFilter)
	filter := spamFilter
	sayHelloWithFilter("Anjing", filter)

}
