package main

import "fmt"

func endApp() {
	fmt.Println("endApp")
}
func runApp(error bool) {
	defer endApp()

	if error {
		panic("Upss error")
	}
}

func main() {
	runApp(true)
}
