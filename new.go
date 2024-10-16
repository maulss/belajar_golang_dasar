package main

import "fmt"

type Address struct {
	Provinsi, Kota, Kabupaten string
}

func main() {
	var alamat1 *Address = new(Address)
	var alamat2 *Address = alamat1
	alamat2.Kota = "Maataram"
	fmt.Println(alamat1)
	fmt.Println(alamat2)
}
