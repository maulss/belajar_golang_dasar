package main

import "fmt"

func logging() {
	fmt.Println("Seelsai memanggil function")
}
func runApplication() {
	defer logging()
	fmt.Println("Run Application")

}

func main() {
	runApplication()
}
