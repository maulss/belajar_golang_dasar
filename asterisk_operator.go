package main

import "fmt"

type Address struct {
	Provinsi, Kota, Kabupaten string
}

func main() {
	address1 := Address{"Jawa timur", "Malang", "subang"}
	address2 := &address1
	address2.Kota = "Bogor"
	fmt.Println(address1)
	fmt.Println(address2)

	*address2 = Address{"Jakarta", "Depok", "SCBD"}
	fmt.Println(address1)
	fmt.Println(address2)
}
