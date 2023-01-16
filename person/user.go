package person

import (
	"mixue/mixue"
	"mixue/ojol"
)

type User struct {
	Name   string
	Saldo  int
	Alamat string
	ojol.Gojek
	mixue.Mixue
}
