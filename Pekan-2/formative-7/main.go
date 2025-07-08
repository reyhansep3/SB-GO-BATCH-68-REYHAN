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

	//soal 3
	soal3()

	//soal 4
	soal4()

	//soal 5
	soal5()
}

// soal 1
func soal1() {
	var user Mahasiswa
	user.Nama = "Reyhan Septri Asta"
	user.NIM = "1301190308"
	user.Usia = 23

	fmt.Println("Nama : ", user.Nama)
	fmt.Println("NIM : ", user.NIM)
	fmt.Println("Usia : ", user.Usia)

}

type Mahasiswa struct {
	Nama string
	NIM  string
	Usia int
}

// soal 2
func soal2() {
	var triangle = segitiga{7, 10}
	var square = persegi{9}
	var rectangle = persegiPanjang{7, 10}

	luasSegitiga := 0.5 * float64(triangle.alas) * float64(triangle.tinggi)
	luasPersegi := square.sisi * square.sisi

	luasPersegiPanjang := rectangle.panjang * rectangle.lebar

	fmt.Println("luas segitiga : ", luasSegitiga)
	fmt.Println("luas persegi : ", luasPersegi)
	fmt.Println("luas persegi panjang : ", luasPersegiPanjang)

}

type segitiga struct {
	alas, tinggi int
}

type persegi struct {
	sisi int
}

type persegiPanjang struct {
	panjang, lebar int
}

// soal 3
func soal3() {
	phoneColors := &[]string{}

	addColors := append(*phoneColors, "Kuning", "Putih", "Biru")

	var phoneUser = phone{"martis", "samsung", 10, addColors}

	fmt.Println(phoneUser)
}

type phone struct {
	name, brand string
	year        int
	colors      []string
}

// soal 4
func soal4() {
	var dataFilm = []movie{}

	tambahDataFilm("LOTR", 120, "action", 1999, &dataFilm)
	tambahDataFilm("avenger", 120, "action", 2019, &dataFilm)
	tambahDataFilm("spiderman", 120, "action", 2004, &dataFilm)
	tambahDataFilm("juon", 120, "horror", 2004, &dataFilm)

	for i, j := range dataFilm {
		fmt.Print(strconv.Itoa(i+1) + ". ")
		fmt.Println("title: ", j.title)
		fmt.Println("   duration ", j.duration)
		fmt.Println("   genre ", j.genre)
		fmt.Println("   year ", j.year)

	}
}

func tambahDataFilm(title string, duration int, genre string, year int, dataFilm *[]movie) {
	*dataFilm = append(*dataFilm, movie{
		title:    title,
		genre:    genre,
		duration: duration,
		year:     year,
	})

}

type movie struct {
	title, genre   string
	duration, year int
}

// soal 5
func soal5() {
	var suaraKucing Kucing
	var suaraAnjing Anjing

	fmt.Println("suara kucing : ", suaraKucing.Suara())
	fmt.Println("suara anjing : ", suaraAnjing.Suara())

}

type Hewan interface {
	Suara() string
}

func (k Kucing) Suara() string {
	return "Meoow"
}

func (a Anjing) Suara() string {
	return "Woof"
}

type Kucing struct {
	suara string
}

type Anjing struct {
	suara string
}
