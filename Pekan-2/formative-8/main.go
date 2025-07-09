package main

import (
	"fmt"
	"formative-8/matematika"
	"math"
	"strconv"
	"strings"
)

func main() {
	//soal 1
	soal1()

	//soal 2
	soal2()

	//soal 3
	soal3()

	//soal 4
	soal4()

	//soal 5
	soal5()
}

// soal 1
func soal1() {
	var bangunDatar hitungBangunDatar
	bangunDatar = segitigaSamaSisi{
		alas:   12,
		tinggi: 15,
	}
	fmt.Println("hasil luas segitiga sama sisi : ", bangunDatar.luas())
	fmt.Println("hasil keliling  segitiga sama sisi: ", bangunDatar.keliling())

	bangunDatar = persegiPanjang{
		lebar:   25,
		panjang: 20,
	}

	fmt.Println("hasil luas persegi panjang : ", bangunDatar.luas())
	fmt.Println("hasil keliling persegi panjang : ", bangunDatar.keliling())

	var bangunRuang hitungBangunRuang
	bangunRuang = tabung{
		jariJari: 25,
		tinggi:   20,
	}

	fmt.Println("hasil volume tabung : ", bangunRuang.volume())
	fmt.Println("hasil luas permukaan tabung : ", bangunRuang.luasPermukaan())

	bangunRuang = balok{
		panjang: 10,
		lebar:   12,
		tinggi:  15,
	}

	fmt.Println("hasil volume balok : ", bangunRuang.volume())
	fmt.Println("hasil luas permukaan balok : ", bangunRuang.luasPermukaan())

}

type segitigaSamaSisi struct {
	alas, tinggi int
}

type persegiPanjang struct {
	panjang, lebar int
}

type tabung struct {
	jariJari, tinggi float64
}

type balok struct {
	panjang, lebar, tinggi int
}

type hitungBangunDatar interface {
	luas() int
	keliling() int
}

type hitungBangunRuang interface {
	volume() float64
	luasPermukaan() float64
}

func (segitiga segitigaSamaSisi) luas() int {
	return (segitiga.alas * segitiga.tinggi) / 2
}

func (segitiga segitigaSamaSisi) keliling() int {
	var sisi = (2 * float64(segitiga.tinggi)) / math.Sqrt(3)
	return 3 * int(math.Round(sisi))
}

func (persegi persegiPanjang) luas() int {
	return persegi.panjang * persegi.lebar
}

func (persegi persegiPanjang) keliling() int {
	return 2 * (persegi.panjang + persegi.lebar)
}

func (t tabung) volume() float64 {
	return math.Phi * float64(t.jariJari*t.jariJari) * float64(t.tinggi)
}

func (t tabung) luasPermukaan() float64 {
	return float64(2 * (math.Phi) * t.jariJari * (t.jariJari * t.tinggi))
}

func (b balok) volume() float64 {
	return float64(b.panjang * b.lebar * b.tinggi)
}

func (b balok) luasPermukaan() float64 {
	return float64(2*b.panjang*b.lebar*b.tinggi + b.lebar*b.tinggi)
}

// soal 2
func soal2() {
	var userPhone detailPhone = phone{
		name:  "Samsung Galaxy Note 20",
		brand: "Samsung Galaxy Note 20",
		year:  2020,
		color: []string{"Mystic Bronze", "Mystic White", "Mystic Black"},
	}

	valueColors := strings.Join(userPhone.warna(), ", ")

	fmt.Println("name : ", userPhone.nama())
	fmt.Println("brand : ", userPhone.produk())
	fmt.Println("year : ", userPhone.tahun())
	fmt.Println("colors : ", valueColors)

}

type phone struct {
	name, brand string
	year        int
	color       []string
}

type detailPhone interface {
	nama() string
	produk() string
	tahun() int
	warna() []string
}

func (p phone) nama() string {
	return p.name
}

func (p phone) produk() string {
	return p.brand
}

func (p phone) tahun() int {
	return p.year
}

func (p phone) warna() []string {
	return p.color
}

// soal 3
func soal3() {

	fmt.Println(luasPersegi(4, true))

	fmt.Println(luasPersegi(8, false))

	fmt.Println(luasPersegi(0, true))

	fmt.Println(luasPersegi(0, false))
}

func luasPersegi(value int, hasil bool) interface{} {
	if hasil && value != 0 {
		luas := value * value
		return "luas persegi dengan sisi " + strconv.Itoa(value) + " cm adalah " + strconv.Itoa(luas)
	} else if !hasil && value != 0 {
		luas := value * value
		return strconv.Itoa(luas)
	} else if hasil && value == 0 {
		return "Maaf anda belum menginput sisi dari persegi"
	} else {
		return nil
	}
}

// soal 4
func soal4() {
	var prefix interface{} = "hasil penjumlahan dari "

	var kumpulanAngkaPertama interface{} = []int{6, 8}

	var kumpulanAngkaKedua interface{} = []int{12, 14}

	var angkaPertama = kumpulanAngkaPertama.([]int)
	var angkaKedua = kumpulanAngkaKedua.([]int)

	hasilKumpulan := angkaPertama[0] + angkaPertama[1] + angkaKedua[0] + angkaKedua[1]
	fmt.Println(hasilKumpulan)

	strAngkaPertama := make([]string, len(angkaPertama))
	strAngkaKedua := make([]string, len(angkaKedua))

	for i, j := range angkaPertama {
		strAngkaPertama[i] = strconv.Itoa(j)
	}

	for x, y := range angkaKedua {
		strAngkaKedua[x] = strconv.Itoa(y)
	}

	joinAngkaPertama := strings.Join(strAngkaPertama, " + ")
	joinAngkaKedua := strings.Join(strAngkaKedua, " + ")

	gabungan := prefix.(string) + joinAngkaPertama + " + " + joinAngkaKedua + " = " + strconv.Itoa(hasilKumpulan)
	fmt.Println(gabungan)

}

// soal 5
func soal5() {
	fmt.Println("hasil tambahan 2 + 2 : ", matematika.Tambah(2, 2))
	fmt.Println("hasil kalian 4 * 5 : ", matematika.Kali(4, 5))

}
