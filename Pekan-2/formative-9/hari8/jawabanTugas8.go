package hari8

import (
	"fmt"
	"formative-9/hari8/matematika"
	"math"
	"strconv"
	"strings"
)

//soal 1

type SegitigaSamaSisi struct {
	Alas, Tinggi int
}

type PersegiPanjang struct {
	Panjang, Lebar int
}

type Tabung struct {
	JariJari, Tinggi float64
}

type Balok struct {
	Panjang, Lebar, Tinggi int
}

type HitungBangunDatar interface {
	Luas() int
	Keliling() int
}

type HitungBangunRuang interface {
	Volume() float64
	LuasPermukaan() float64
}

func (segitiga SegitigaSamaSisi) Luas() int {
	return (segitiga.Alas * segitiga.Tinggi) / 2
}

func (segitiga SegitigaSamaSisi) Keliling() int {
	var sisi = (2 * float64(segitiga.Tinggi)) / math.Sqrt(3)
	return 3 * int(math.Round(sisi))
}

func (persegi PersegiPanjang) Luas() int {
	return persegi.Panjang * persegi.Lebar
}

func (persegi PersegiPanjang) Keliling() int {
	return 2 * (persegi.Panjang + persegi.Lebar)
}

func (t Tabung) Volume() float64 {
	return math.Phi * float64(t.JariJari*t.JariJari) * float64(t.Tinggi)
}

func (t Tabung) LuasPermukaan() float64 {
	return float64(2 * (math.Phi) * t.JariJari * (t.JariJari * t.Tinggi))
}

func (b Balok) Volume() float64 {
	return float64(b.Panjang * b.Lebar * b.Tinggi)
}

func (b Balok) LuasPermukaan() float64 {
	return float64(2*b.Panjang*b.Lebar*b.Tinggi + b.Lebar*b.Tinggi)
}

// soal 2
type Phone struct {
	Name, Brand string
	Year        int
	Color       []string
}

type DetailPhone interface {
	Nama() string
	Produk() string
	Tahun() int
	Warna() []string
}

func (p Phone) Nama() string {
	return p.Name
}

func (p Phone) Produk() string {
	return p.Brand
}

func (p Phone) Tahun() int {
	return p.Year
}

func (p Phone) Warna() []string {
	return p.Color
}

// soal 3
func LuasPersegi(value int, hasil bool) interface{} {
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
func Soal4() {
	var prefix interface{} = "hasil penjumlahan dari "

	var kumpulanAngkaPertama interface{} = []int{6, 8}

	var kumpulanAngkaKedua interface{} = []int{12, 14}

	var angkaPertama = kumpulanAngkaPertama.([]int)
	var angkaKedua = kumpulanAngkaKedua.([]int)

	hasilKumpulan := angkaPertama[0] + angkaPertama[1] + angkaKedua[0] + angkaKedua[1]

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
func Soal5() {
	fmt.Println("hasil tambahan 2 + 2 : ", matematika.Tambah(2, 2))
	fmt.Println("hasil kalian 4 * 5 : ", matematika.Kali(4, 5))
}
