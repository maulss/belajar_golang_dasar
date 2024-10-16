package main

import "fmt"

type Address struct {
	Provinsi, Kota, Kabupaten string
}

func changeAddresToIndonesia(addr *Address) {
	addr.Provinsi = "jakarta"
}
func main() {
	address := Address{"NTb", "Mataram", "dsandubaya"}
	changeAddresToIndonesia(&address)
	fmt.Println(address)

}
