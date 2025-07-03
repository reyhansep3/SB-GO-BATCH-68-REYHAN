package main

import (
	"fmt"
	"strconv"
)

func main() {
	//soal 1
	var hasilPiramid = ""
	var i = 0
	var j = 0
	var jumlahPagar = 7

	for i < jumlahPagar {
		for j < jumlahPagar && j == i {
			hasilPiramid += "#"
			j++
		}
		hasilPiramid += ""
		i++
		fmt.Println(hasilPiramid)
	}

	//soal 2
	var kalimat = [...]string{"aku", "dan", "saya", "sangat", "senang", "belajar", "golang"}

	var hasil = [5]string{
		kalimat[2],
		kalimat[3],
		kalimat[4],
		kalimat[5],
		kalimat[6],
	}

	fmt.Println(hasil)

	//soal 3
	var sayuran = []string{}

	// Bayam
	// Buncis
	// Kangkung
	// Kubis
	// Seledri
	// Tauge
	// Timun

	sayuran = append(sayuran, "Bayam")
	sayuran = append(sayuran, "Buncis")
	sayuran = append(sayuran, "Kangkung")
	sayuran = append(sayuran, "Kubis")
	sayuran = append(sayuran, "Seledri")
	sayuran = append(sayuran, "Tauge")
	sayuran = append(sayuran, "Timun")

	fmt.Println(sayuran)

	for i := 0; i < len(sayuran); i++ {
		var jadikanStr = strconv.Itoa(i + 1)
		sayuran[i] = jadikanStr + ". " + sayuran[i]
		fmt.Println(sayuran[i])
	}

	//soal 4
	var satuan = map[string]int{
		"panjang": 7,
		"lebar":   4,
		"tinggi":  6,
	}
	fmt.Println("panjang =", satuan["panjang"])
	fmt.Println("lebar =", satuan["lebar"])
	fmt.Println("tinggi =", satuan["tinggi"])
	fmt.Println("volume balok =", satuan["panjang"]*satuan["lebar"]*satuan["tinggi"])

	//soal 5
	panjang := 12
	lebar := 4
	tinggi := 8

	luas := luasPersegiPanjang(panjang, lebar)
	keliling := kelilingPersegiPanjang(panjang, lebar)
	volume := volumeBalok(panjang, lebar, tinggi)

	fmt.Println(luas)
	fmt.Println(keliling)
	fmt.Println(volume)

}

// function untuk soal no 5
func luasPersegiPanjang(nilaiPanjang int, nilaiLebar int) int {
	return nilaiPanjang * nilaiLebar
}

func kelilingPersegiPanjang(nilaiPanjang int, nilaiLebar int) int {
	return 2 * (nilaiPanjang + nilaiLebar)
}

func volumeBalok(nilaiPanjang int, nilaiLebar int, nilaiTinggi int) int {
	return nilaiPanjang * nilaiLebar * nilaiTinggi
}
