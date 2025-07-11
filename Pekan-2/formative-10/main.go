package main

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"sync"
	"time"
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
	soal4()
	fmt.Println("")

	//soal 5
	soal5()
}

func soal1() {
	defer tampilkan("Golang Backend Development", 2021)
}

func tampilkan(kalimat string, tahun int) {
	fmt.Print(kalimat + " ")
	fmt.Println(tahun)
}

func soal2() {
	fmt.Println(kelilingSegitigaSamaSisi(4, true))
	fmt.Println(kelilingSegitigaSamaSisi(8, false))
	fmt.Println(kelilingSegitigaSamaSisi(0, true))
	fmt.Println(kelilingSegitigaSamaSisi(0, false))
}

func kelilingSegitigaSamaSisi(nilai int, jenis bool) (string, error) {

	if nilai != 0 {
		if jenis {
			hasilKeliling := nilai * 3
			var kalimat = "keliling segitiga sama sisinya dengan sisi " + strconv.Itoa(nilai) + " cm adalah " + strconv.Itoa(hasilKeliling)
			return kalimat, errors.New("")
		} else {
			hasilKeliling := nilai * 3
			return strconv.Itoa(hasilKeliling), errors.New("")
		}
	} else {
		if jenis {
			return "", errors.New("maaf anda belum menginput sisi dari segitiga sama sisi")
		} else {
			defer recoverPanic()
			panic("Maaf anda belum menginput sisi dari segitiga sama sisi")
		}
	}
}

func recoverPanic() {
	message := recover()
	fmt.Print("", message)
}

// soal 3
func soal3() {
	angka := 1

	defer cetakAngka(&angka)

	tambahAngka(7, &angka)

	tambahAngka(6, &angka)

	tambahAngka(-1, &angka)

	tambahAngka(9, &angka)
}

func tambahAngka(value int, valueAngka *int) {
	*valueAngka += value
}

func cetakAngka(valueAngka *int) {
	fmt.Println("Total angka :", *valueAngka)
}

// soal 4
func soal4() {
	var phones = []string{}
	tambahData(&phones)
}

func tambahData(data *[]string) {
	*data = append(*data, "Xiaomi", "Asus", "IPhone", "Samsung", "Oppo", "Realme", "Vivo")
	sort.Strings(*data)

	for i, j := range *data {
		time.Sleep(time.Second * 1)
		fmt.Print(strconv.Itoa(i+1) + ". ")
		fmt.Println(j)
	}
}

// soal 5
func soal5() {
	var wg sync.WaitGroup

	var phones = []string{"Xiaomi", "Asus", "Iphone", "Samsung", "Oppo", "Realme", "Vivo"}
	sort.Strings(phones)

	wg.Add(1)
	go tampilkanData(&phones, &wg)

	wg.Wait()
}

func tampilkanData(data *[]string, wg *sync.WaitGroup) {

	for i, j := range *data {

		fmt.Print(strconv.Itoa(i+1) + ". ")
		fmt.Println(j)
		time.Sleep(1 * time.Second)
	}
	wg.Done()
}
