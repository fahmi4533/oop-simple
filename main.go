package main

import (
	"fmt"
	"mixue/person"
)

func main() {

	user1 := person.User{
		Name:   "AFDAL GOZALI",
		Saldo:  5000000,
		Alamat: "Marpoyan",
	}

	user1.CreatPesan = map[string]int{
		"es krim kocok":   15,
		"es krim basah":   20,
		"es krim pelumas": 30,
		"nasi kocok":      24,
		"ayam":            27,
		"teh es":          30,
	}

	user1.Gojek.PesanGoFood(user1.Name)
	user1.Gojek.TarifOngkir(user1.Alamat)
	user1.Gojek.TotalKeseluruhan()
	user1.Saldo = user1.Gojek.Bayar(user1.Saldo)
	fmt.Println("SISA SALDO ANDA :", user1.Saldo)
	fmt.Println("=============================================")

}
