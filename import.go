package main

import (
	"belajar_golang_dasar/helper"
	"fmt"
)

func main() {
	result := helper.SayHello("Maulana")
	fmt.Println(result)

	fmt.Println(helper.Application)
	fmt.Println(helper.SayHello("kocak"))
}
