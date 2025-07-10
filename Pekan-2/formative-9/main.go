package main

import (
	"fmt"
	"formative-9/hari8"
	"strings"
)

func main() {
	//soal 1
	soal1()

	fmt.Println("")
	//soal 2
	soal2()

	fmt.Println("")
	//soal 3
	soal3()

	fmt.Println("")
	//soal 4
	hari8.Soal4()

	fmt.Println("")
	//soal 5
	hari8.Soal5()
}

// soal 1
func soal1() {
	var bangunDatar hari8.HitungBangunDatar
	bangunDatar = hari8.SegitigaSamaSisi{
		Alas:   12,
		Tinggi: 15,
	}
	fmt.Println("hasil luas segitiga sama sisi : ", bangunDatar.Luas())
	fmt.Println("hasil keliling  segitiga sama sisi: ", bangunDatar.Keliling())

	bangunDatar = hari8.PersegiPanjang{
		Lebar:   25,
		Panjang: 20,
	}

	fmt.Println("hasil luas persegi panjang : ", bangunDatar.Luas())
	fmt.Println("hasil keliling persegi panjang : ", bangunDatar.Keliling())

	var bangunRuang hari8.HitungBangunRuang
	bangunRuang = hari8.Tabung{
		JariJari: 25,
		Tinggi:   20,
	}

	fmt.Println("hasil volume tabung : ", bangunRuang.Volume())
	fmt.Println("hasil luas permukaan tabung : ", bangunRuang.LuasPermukaan())

	bangunRuang = hari8.Balok{
		Panjang: 10,
		Lebar:   12,
		Tinggi:  15,
	}

	fmt.Println("hasil volume balok : ", bangunRuang.Volume())
	fmt.Println("hasil luas permukaan balok : ", bangunRuang.LuasPermukaan())

}

// soal 2
func soal2() {
	var userPhone hari8.DetailPhone = hari8.Phone{
		Name:  "Samsung Galaxy Note 20",
		Brand: "Samsung Galaxy Note 20",
		Year:  2020,
		Color: []string{"Mystic Bronze", "Mystic White", "Mystic Black"},
	}

	valueColors := strings.Join(userPhone.Warna(), ", ")

	fmt.Println("name : ", userPhone.Nama())
	fmt.Println("brand : ", userPhone.Produk())
	fmt.Println("year : ", userPhone.Tahun())
	fmt.Println("colors : ", valueColors)

}

// soal 3
func soal3() {

	fmt.Println(hari8.LuasPersegi(4, true))

	fmt.Println(hari8.LuasPersegi(8, false))

	fmt.Println(hari8.LuasPersegi(0, true))

	fmt.Println(hari8.LuasPersegi(0, false))
}
