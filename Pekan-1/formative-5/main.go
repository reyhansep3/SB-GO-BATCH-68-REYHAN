package main

import (
	"fmt"
	"math"
)

func main() {
	//soal 1
	john := introduce("John", "laki-laki", "penulis", "30")
	sarah := introduce("Sarah", "perempuan", "model", "28")
	fmt.Println(john)
	fmt.Println(sarah)

	//soal 2
	var buah = []string{"semangka", "jeruk", "melon", "pepaya"}
	var buahFavoritJohn = buahFavorit("John", buah...)
	fmt.Println(buahFavoritJohn)

	//soal 3
	var dataFilm = []map[string]string{}

	var tambahDataFilm = func(name string, hour string, genre string, year string) {
		dataFilm = append(dataFilm, map[string]string{
			"genre": genre,
			"jam":   hour,
			"tahun": year,
			"title": name,
		})
	}
	tambahDataFilm("LOTR", "2 jam", "action", "1999")
	tambahDataFilm("avenger", "2 jam", "action", "2019")
	tambahDataFilm("spiderman", "2 jam", "action", "2004")
	tambahDataFilm("juon", "2 jam", "horror", "2004")

	for _, item := range dataFilm {
		fmt.Println(item)
	}

	//soal 4
	factorial(5)
	factorial(7)

	//soal 5
	var jarijari float64 = 7
	var luasLingkaran float64
	var kelilingLingkaran float64
	perhitungan(&luasLingkaran, &kelilingLingkaran, jarijari)

}

// soal 1
func introduce(nama string, gender string, job string, age string) (hasil string) {
	var panggilan = ""
	if gender == "laki-laki" {
		panggilan = "Pak"
	} else {
		panggilan = "Bu"
	}
	hasil = panggilan + " " + nama + " " + "adalah seorang" + " " + job + " " + "yang berusia" + " " + age + " " + "tahun"
	return

}

// soal 2
func buahFavorit(name string, nameBuah ...string) string {
	buahFavorit := ""
	for i, name := range nameBuah {
		if i != len(nameBuah)-1 {
			buahFavorit += name + " "

		} else {
			buahFavorit += name
		}
	}

	hasil := "halo nama saya " + name + " dan buah favorit saya adalah " + buahFavorit + ""
	return hasil
}

// soal 4
func factorial(angka int) {
	var hasil = 1
	for angka >= 1 {
		hasil *= angka
		angka--
	}
	fmt.Println(hasil)
}

// soal 5
func perhitungan(luas *float64, keliling *float64, jarijari float64) {

	*luas = math.Pi * (jarijari * jarijari)
	*keliling = 2 * math.Pi * jarijari

	fmt.Println("hasil luas lingkaran : ", *luas)
	fmt.Println("hasil keliling lingkaran : ", *keliling)

}
