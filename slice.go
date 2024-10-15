package main

import "fmt"

func main() {
	var newSlice []string = make([]string, 2, 5)
	newSlice[0] = "maulana"
	newSlice[1] = "givari"
	//newSlice[2] = "givari" error, harus menggunakkan append

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice2 := append(newSlice, "givari")
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))

	newSlice2[0] = "muhammad"
	fmt.Println(newSlice2)
	fmt.Println(newSlice)
}
