package main

import "fmt"

func endApps() {
	fmt.Println("endApp")
	message := recover()
	fmt.Println("terjadi panic", message)
}
func runApps(error bool) {
	defer endApps()

	if error {
		panic("Upss error")
	}

}

func main() {
	runApps(true)
	fmt.Println("aplikasi tetap dijalakna")
}
