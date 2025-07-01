package main

import (
	"fmt"
	"strconv"
)

func main() {
	//soal 1
	fmt.Println("Bootcamp Digital Skill Sanbercode Golang")

	//soal 2
	halo := "Halo Dunia"
	halo = "Golang Dunia"
	fmt.Println(halo)

	//soal 3
	var kataPertama = "saya"
	var kataKedua = "senang"
	var kataKetiga = "belajar"
	var kataKeempat = "golang"

	var sentence = kataPertama + " " + kataKedua + " " + kataKetiga + " " + kataKeempat

	fmt.Println(sentence)

	//soal 4
	var angkaPertama = "8"
	var angkaKedua = "5"
	var angkaKetiga = "6"
	var angkaKeempat = "7"

	numPertama, err := strconv.Atoi(angkaPertama)

	numKedua, err := strconv.Atoi(angkaKedua)

	numKetiga, err := strconv.Atoi(angkaKetiga)

	numKeempat, err := strconv.Atoi(angkaKeempat)

	if err == nil {
		var hasil = numPertama + numKedua + numKetiga + numKeempat

		fmt.Println("hasil perjumlahan : ", hasil)
	}

	//soal 5

	var panjangPersegiPanjang string = "8"
	var lebarPersegiPanjang string = "5"

	var alasSegitiga string = "6"
	var tinggiSegitiga string = "7"

	intPanjangPersegiPanjang, err := strconv.Atoi(panjangPersegiPanjang)
	intLebarPersegiPanjang, err := strconv.Atoi(lebarPersegiPanjang)
	intAlasSegitiga, err := strconv.Atoi(alasSegitiga)
	intTinggiSegitiga, err := strconv.Atoi(tinggiSegitiga)

	var luasPersegiPanjang = intPanjangPersegiPanjang * intLebarPersegiPanjang
	var kelilingPersegiPanjang = 2 * (intPanjangPersegiPanjang * intLebarPersegiPanjang)
	var luasSegitiga = (2 * intTinggiSegitiga) / intAlasSegitiga

	fmt.Println("luas persegi panjang : ", luasPersegiPanjang)
	fmt.Println("keliling persegi panjang : ", kelilingPersegiPanjang)
	fmt.Println("luas segitiga : ", luasSegitiga)

}
