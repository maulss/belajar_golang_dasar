package main

import "fmt"

func main() {
	counter := 1

	for counter < 10 {
		fmt.Println("counter ", counter)
		counter++
	}

	for number := 1; number < 10; number++ {
		fmt.Println("number ke", number)
	}

	names := []string{"Muhammad", "Maulana", "Givari"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	for index, name := range names {
		fmt.Println(index, name)
	}

	for _, name := range names {
		fmt.Println(name)
	}
}
