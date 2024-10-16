package main

import "fmt"

func factorialLopp(x int) int {
	result := 1
	for i := x; i > 0; i-- {
		result *= i
	}
	return result
}

func factorialRecursive(x int) int {
	if x == 1 {
		return 1
	} else {
		return x * factorialRecursive(x-1)
	}
}

func main() {
	fmt.Println(factorialLopp(10))
	fmt.Println(factorialRecursive(10))
}
