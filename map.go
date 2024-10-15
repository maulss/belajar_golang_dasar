package main

import "fmt"

func main() {
	person := map[string]string{
		"name":   "maulana",
		"umur":   "12",
		"alamat": "test",
	}

	fmt.Println(person["name"])
	fmt.Println(person["umur"])
	fmt.Println(person["alamat"])

	book := map[string]string{}

	book["judul"] = "judul test"
	book["tahun terbit"] = "2012"
	book["penerbit"] = "yahoo"

	fmt.Println(book)
	book["judul"] = "judul baru"
	fmt.Println(book)

}
