package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//soal 1
	kalimat := "halo halo bandung"
	angka := 2021

	kalimatReplaced := strings.Replace(kalimat, "halo", "hi", 2)
	angkaStr := strconv.Itoa(angka)

	output := "\"" + kalimatReplaced + "\" - " + angkaStr

	fmt.Println(output)

	//soal 2
	var nilaiJohn = 80
	var nilaiDoe = 50

	if nilaiJohn >= 80 {
		fmt.Println("Nilai jhon : A")
		if nilaiDoe >= 80 {
			fmt.Println("Nilai doe : A")
		} else if nilaiDoe >= 70 && nilaiDoe < 80 {
			fmt.Println("Nilai doe : B")
		} else if nilaiDoe >= 60 && nilaiDoe < 70 {
			fmt.Println("Nilai doe : C")
		} else if nilaiDoe >= 50 && nilaiDoe < 60 {
			fmt.Println("Nilai doe : D")
		} else {
			fmt.Println("Nilai doe : E")
		}
	} else if nilaiJohn >= 70 && nilaiJohn < 80 {
		fmt.Println("Nilai jhon : B")
		if nilaiDoe >= 80 {
			fmt.Println("Nilai doe : A")
		} else if nilaiDoe >= 70 && nilaiDoe < 80 {
			fmt.Println("Nilai doe : B")
		} else if nilaiDoe >= 60 && nilaiDoe < 70 {
			fmt.Println("Nilai doe : C")
		} else if nilaiDoe >= 50 && nilaiDoe < 60 {
			fmt.Println("Nilai doe : D")
		} else {
			fmt.Println("Nilai doe : E")
		}
	} else if nilaiJohn >= 60 && nilaiJohn < 70 {
		fmt.Println("Nilai jhon : C")
		if nilaiDoe >= 80 {
			fmt.Println("Nilai doe : A")
		} else if nilaiDoe >= 70 && nilaiDoe < 80 {
			fmt.Println("Nilai doe : B")
		} else if nilaiDoe >= 60 && nilaiDoe < 70 {
			fmt.Println("Nilai doe : C")
		} else if nilaiDoe >= 50 && nilaiDoe < 60 {
			fmt.Println("Nilai doe : D")
		} else {
			fmt.Println("Nilai doe : E")
		}
	} else if nilaiJohn >= 50 && nilaiJohn < 60 {
		fmt.Println("Nilai jhon : D")
		if nilaiDoe >= 80 {
			fmt.Println("Nilai doe : A")
		} else if nilaiDoe >= 70 && nilaiDoe < 80 {
			fmt.Println("Nilai doe : B")
		} else if nilaiDoe >= 60 && nilaiDoe < 70 {
			fmt.Println("Nilai doe : C")
		} else if nilaiDoe >= 50 && nilaiDoe < 60 {
			fmt.Println("Nilai doe : D")
		} else {
			fmt.Println("Nilai doe : E")
		}
	} else {
		fmt.Println("Nilai jhon : E")
		if nilaiDoe >= 80 {
			fmt.Println("Nilai doe : A")
		} else if nilaiDoe >= 70 && nilaiDoe < 80 {
			fmt.Println("Nilai doe : B")
		} else if nilaiDoe >= 60 && nilaiDoe < 70 {
			fmt.Println("Nilai doe : C")
		} else if nilaiDoe >= 50 && nilaiDoe < 60 {
			fmt.Println("Nilai doe : D")
		} else {
			fmt.Println("Nilai doe : E")
		}
	}

	//soal 3
	var tanggal = 15
	var bulan = 9
	var tahun = 2001

	namaBulan := strconv.Itoa(bulan)
	namaTanggal := strconv.Itoa(tanggal)
	namaTahun := strconv.Itoa(tahun)

	switch bulan {
	case 1:
		namaBulan = "Januari"
	case 2:
		namaBulan = "Februari"
	case 3:
		namaBulan = "Maret"
	case 4:
		namaBulan = "April"
	case 5:
		namaBulan = "Mei"
	case 6:
		namaBulan = "Juni"
	case 7:
		namaBulan = "Juli"
	case 8:
		namaBulan = "Agustus"
	case 9:
		namaBulan = "September"
	case 10:
		namaBulan = "Oktober"
	case 11:
		namaBulan = "November"
	case 12:
		namaBulan = "desember"
	}

	fmt.Println(namaTanggal + " " + namaBulan + " " + namaTahun)

	//soal 4
	var kelahiran = 2001

	if kelahiran >= 1944 && kelahiran <= 1964 {
		fmt.Println("saya generasi Boomer")
	} else if kelahiran >= 1965 && kelahiran <= 1979 {
		fmt.Println("saya generasi Generasi X")
	} else if kelahiran >= 1980 && kelahiran <= 1994 {
		fmt.Println("saya generasi Generasi Y (Millenials)")
	} else if kelahiran >= 1995 && kelahiran <= 2015 {
		fmt.Println("saya generasi Generasi Z")
	}

	//soal 5
	var hitungan = 0
	var angkaMasukan = 20

	for hitungan != angkaMasukan {
		if hitungan%2 == 0 {
			fmt.Print(hitungan)
			fmt.Println(" - berkualitas")
		} else if hitungan%3 == 0 && hitungan%2 == 1 {
			fmt.Print(hitungan)
			fmt.Println(" - I Love Coding ")
		} else if hitungan%2 == 1 {
			fmt.Print(hitungan)
			fmt.Println(" - santai")
		}
		hitungan = hitungan + 1
	}
}
