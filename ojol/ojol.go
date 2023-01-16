package ojol

import (
	"fmt"
	"mixue/mixue"
	"time"
)

type Gojek struct {
	Tarif      int
	PesanGofod int
	Total      int
	CreatPesan map[string]int
}

func (gojek *Gojek) PesanGoFood(name string) {
	mixue1 := mixue.Mixue{}
	mixue1.Menue()
	mixue1.LokasiMixue()

	Pesanan1 := mixue.PesananCustomer{
		NameUser: name,
		Pesanan:  gojek.CreatPesan,
	}

	mixue.InputPesanan(Pesanan1.NameUser, Pesanan1.Pesanan)
	mixue1.StorePesanan(Pesanan1.Pesanan)
	gojek.Total = mixue1.TotalGojek()
}

func (gojek *Gojek) TarifOngkir(alamat string) {
	pekanbaru := []string{"Panam", "ArifinAhmad", "Sudirman", "Marpoyan", "Kubang", "PasirPutih"}
	mixue1 := mixue.Mixue{}

	tarif := 8000
	km := 0
	j := 0
	k := 0
	for i, s := range pekanbaru {
		if s == alamat {
			j = i + 1
		}
		if s == mixue1.Lokasi {
			k = i + 1
		}
	}

	km = j - k
	if km < -km {
		km = -km
	}

	for i := 0; i < km; i++ {
		gojek.Tarif += tarif
	}
	fmt.Println("=============================================")
	time.Sleep(2 * time.Second)
	fmt.Println("TARIF ONGKIR PESANAN ANDA :", gojek.Tarif)
}

func (gojek *Gojek) TotalKeseluruhan() {
	fmt.Println("=============================================")
	time.Sleep(2 * time.Second)
	fmt.Println("TOTAL BELANJAAN ANDA DI Mixue :", gojek.Total)
	fmt.Println("=============================================")
	fmt.Println("TOTAL KESELURUHAN BELANJAAN ANDA : ", gojek.Tarif+gojek.Total)
}

func (gojek *Gojek) Bayar(saldo int) int {
	if saldo < gojek.Total+gojek.Tarif {
		fmt.Println("=============================================")
		time.Sleep(2 * time.Second)
		fmt.Println("--MAAF SALDO ANDA TIDAN MENCUKUPI--")
		fmt.Println("=============================================")
	} else {
		saldo -= gojek.Total + gojek.Tarif
		fmt.Println("=============================================")
		time.Sleep(2 * time.Second)
		fmt.Println("--PEMBAYARAN BERHASIL--")
		fmt.Println("=============================================")
	}

	return saldo
}
