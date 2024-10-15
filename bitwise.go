package main

import "fmt"

func main() {
	a := 10
	b := 20

	fmt.Println("a & b", a&b)
	fmt.Println("a | b", a|b)
	fmt.Println("a ^ b", a^b)
	fmt.Println("a << 1", a<<11)
	fmt.Println("a >> 1", a>>1)
}
