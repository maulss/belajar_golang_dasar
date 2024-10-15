package main

import "fmt"

func main() {
	a := 20
	b := 10

	fmt.Println("a + b =", a+b)
	fmt.Println("a - b =", a-b)
	fmt.Println("a * b =", a*b)
	fmt.Println("a / b =", a/b)
	fmt.Println("a % b =", a%b)

	var i = 10
	i += 10
	fmt.Println("i : ", i)
}
