package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "A"
	names[1] = "B"
	names[2] = "C"

	fmt.Println(names)

	var values = [...]int{
		1, 2, 3, 8, 10, 20, 6,
	}

	fmt.Println(values)
	fmt.Println("panjang array", len(values))
	fmt.Println(values[0])
	values[2] = 200
	fmt.Println(values)
}
