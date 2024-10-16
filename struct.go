package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "My name is ", customer.Name)
}

func main() {
	var customer Customer
	fmt.Println(customer)
	customer.Name = "Muhammad Maulana Givari"
	customer.Address = "Jln test"
	customer.Age = 18
	fmt.Println(customer)

	joko := Customer{
		Name:    "joko",
		Age:     18,
		Address: "Test",
	}
	fmt.Println(joko)

	buddy := Customer{"Budi", "test", 18}
	fmt.Println(buddy)

	buddy.sayHello("James bond")

}
