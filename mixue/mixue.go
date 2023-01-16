package mixue

import (
	"fmt"
	"time"
)

type Mixue struct {
	Lokasi     string
	Menu       map[string]int
	Pembayaran struct {
		Total int
	}
	Pesanan struct {
		Note  []string
		Order map[string]int
	}
	*PesananCustomer
}

type PesananCustomer struct {
	NameUser string
	Pesanan  map[string]int
}

func (mixue *Mixue) LokasiMixue() {
	mixue.Lokasi = "Panam"
}

func (mixue *Mixue) StorePesanan(arr1 map[string]int) {
	mixue.Pesanan.Order = arr1

	for k := range arr1 {
		mixue.Pesanan.Note = append(mixue.Pesanan.Note, k)
	}
	fmt.Println("=============================================")
	time.Sleep(2 * time.Second)
	fmt.Println("--BERHASIL MEMESAN MENU--")
}

func InputPesanan(name string, pesanan map[string]int) {
	time.Sleep(2 * time.Second)
	fmt.Println("--MEMBUAT PESANAN--")
	time.Sleep(2 * time.Second)
	fmt.Println("Pemilik Pesanan :", name)
	time.Sleep(1 * time.Second)

	fmt.Println("----ORDERAN----")
	for k, v := range pesanan {
		fmt.Println(k, ":", v)
		time.Sleep(1 * time.Second)
	}
}

func (mixue *Mixue) LihatMenuMixue() {
	fmt.Println("=============================================")
	fmt.Println("--Menu Mixue--")
	for k, v := range mixue.Menu {
		fmt.Println(k, ":", v)
	}
	fmt.Println("=============================================")
}

func (mixue *Mixue) TotalGojek() int {
	sumTotal := 0

	for _, s := range mixue.Pesanan.Note {
		cek := mixue.Pesanan.Order[s]
		for i := 0; i < cek; i++ {
			sumTotal += mixue.Menu[s]
		}
	}

	mixue.Pembayaran.Total = sumTotal
	return mixue.Pembayaran.Total
}

func (mixue *Mixue) Total() {
	sumTotal := 0

	for _, s := range mixue.Pesanan.Note {
		cek := mixue.Pesanan.Order[s]
		for i := 0; i < cek; i++ {
			sumTotal += mixue.Menu[s]
		}
	}

	mixue.Pembayaran.Total = sumTotal
	time.Sleep(3 * time.Second)
	fmt.Println("Total Seluruh pesanan anda :", sumTotal, "Ribu")
}

func (mixue *Mixue) Bayar(val int) int {
	if val < mixue.Pembayaran.Total {
		fmt.Println("Maaf Saldo Anda Tidak Mencukupi")
	} else {
		val -= mixue.Pembayaran.Total
		fmt.Println("PEMBAYARAN BERHASIL")
		time.Sleep(2 * time.Second)
	}
	return val
}

func (mixue *Mixue) Menue() {
	listMenu := map[string]int{
		"es krim kocok":   15000,
		"es krim basah":   20000,
		"es krim pelumas": 30000,
		"nasi kocok":      25000,
		"ayam":            23000,
		"teh es":          5000,
	}
	mixue.Menu = listMenu
}
