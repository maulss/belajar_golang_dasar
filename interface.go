package main

import "fmt"

type HasName interface {
	GetName() string
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func SayHello(value HasName) {

	fmt.Println("Hello", value.GetName())
}

func main() {
	person := Person{Name: "Maulana"}
	animal := Animal{"Test"}
	fmt.Println(person.Name)
	SayHello(person)
	SayHello(animal)
}
