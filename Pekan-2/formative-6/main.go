package main

import (
	"fmt"
	"strconv"
)

func main() {
	//soal 1
	soal1()

	//soal 2
	soal2()

	//soal3
	soal3()

	//soal 4
	soal4()

	//soal 5
	soal5()
}

func soal1() {
	var sentence string
	introduce(&sentence, "jhon", "laki-laki", "penulis", "30")
	fmt.Println(sentence)

	introduce(&sentence, "sarah", "perempuan", "model", "28")
	fmt.Println(sentence)
}

func introduce(sentence *string, name string, gender string, job string, age string) {
	var title string
	if gender == "perempuan" {
		title = "Bu"
	} else {
		title = "Pak"
	}
	*sentence = title + " " + name + " " + "adalah seorang" + " " + job + " " + "yang berusia" + " " + age + " " + "tahun"
}

func soal2() {
	var buah []string

	tambahData(&buah)

	for i, semuaData := range buah {
		fmt.Println(strconv.Itoa(i+1) + "." + " " + semuaData)
	}

}

func tambahData(array *[]string) {
	*array = append(*array, "jeruk", "Semangka", "Mangga", "Strawberry", "Durian", "Manggis", "Alpukat")
}

func soal3() {
	var dataFilm = []map[string]string{}

	tambahDataFilm("LOTR", "2 jam", "action", "1999", &dataFilm)
	tambahDataFilm("avenger", "2 jam", "action", "2019", &dataFilm)
	tambahDataFilm("spiderman", "2 jam", "action", "2004", &dataFilm)
	tambahDataFilm("juon", "2 jam", "horror", "2004", &dataFilm)

	hasilTambahFilm(&dataFilm)
}

func tambahDataFilm(name string, hours string, genre string, year string, film *[]map[string]string) {
	dataFilm := map[string]string{
		"title":    name,
		"duration": hours,
		"genre":    genre,
		"year":     year,
	}
	*film = append(*film, dataFilm)
}

func hasilTambahFilm(film *[]map[string]string) {
	for i, semuaFilm := range *film {
		fmt.Print(strconv.Itoa(i+1) + " ")
		fmt.Println("title:", semuaFilm["title"])
		fmt.Println("  duration:", semuaFilm["duration"])
		fmt.Println("  genre:", semuaFilm["genre"])
		fmt.Println("  year:", semuaFilm["year"])
	}
}

func soal4() {
	var array = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	ubahAngka(&array)
}

func ubahAngka(angka *[9]int) {
	fmt.Println("sebelum : ", *angka)
	var hasil = angka
	for i, semuaAngka := range hasil {
		fmt.Print(strconv.Itoa(i+1) + "." + " " + strconv.Itoa(semuaAngka))
		if semuaAngka%2 != 0 {
			semuaAngka += 1
			fmt.Print(" + 1")
		} else {
			fmt.Print("    ")
		}
		semuaAngka *= 2
		fmt.Println(" * 2 :", semuaAngka)

		hasil[i] = semuaAngka
	}
	fmt.Println("setelah : ", *hasil)
}

func soal5() {
	var hasil buah

	var namaBuah = []string{"Nanas", "Jeruk", "Semangka", "Pisang"}
	var warnaBuah = []string{"Kuning", "Oranye", "Hijau & Merah", "Kuning"}
	var adaBijiBuah = []bool{false, true, false, true}
	var hargaBuah = []int{9000, 8000, 10000, 5000}

	for i := 0; i < len(namaBuah); i++ {
		hasil.nama = namaBuah[i]
		hasil.warna = warnaBuah[i]
		hasil.adaBijinya = adaBijiBuah[i]
		hasil.harga = hargaBuah[i]

		fmt.Print(hasil.nama)
		fmt.Print(" ")
		fmt.Print(hasil.warna)
		fmt.Print(" ")
		fmt.Print(hasil.adaBijinya)
		fmt.Print(" ")
		fmt.Print(hasil.harga)
		fmt.Println("")
	}

}

type buah struct {
	nama       string
	warna      string
	adaBijinya bool
	harga      int
}
