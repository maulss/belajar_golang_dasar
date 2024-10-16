package main

import "fmt"

func random() any {
	return "OK"
}

func main() {
	//random := random()
	//var resultString string = random.(string)
	//fmt.Println(resultString)
	result := random()
	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Int", value)

	default:
		fmt.Println("Unknow Type")
	}
}
