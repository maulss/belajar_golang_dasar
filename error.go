package main

import (
	"errors"
	"fmt"
)

func pembagian(nilai, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("pembagi is zero")
	} else {
		hasil := nilai / pembagi
		return hasil, nil
	}
}
func main() {
	hasil, err := pembagian(100, 0)
	if err == nil {
		fmt.Println("Hasil : ", hasil)
	} else {
		fmt.Println("erros", err.Error())
	}
}
