package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	//soal 1
	http.Handle("/indeks", Auth(http.HandlerFunc(nilaiMahasiswa)))

	//soal 2
	http.Handle("/getIndeks", Auth(http.HandlerFunc(tampilNilai)))

	fmt.Println("server running at http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}

type NilaiMahasiswa struct {
	Nama        string `json:"nama"`
	MataKuliah  string `json:"mataKuliah"`
	IndeksNilai string `json:"indeksNilai"`
	Nilai       uint   `json:"nilai"`
	ID          uint   `json:"id"`
}

// soal 2
var semuaDataMahasiswa []NilaiMahasiswa

// soal 1
func nilaiMahasiswa(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var nilaiMahasiswa NilaiMahasiswa
		if r.Header.Get("Content-Type") == "application/json" {
			decodeJson := json.NewDecoder(r.Body)
			if err := decodeJson.Decode(&nilaiMahasiswa); err != nil {
				log.Fatal(err.Error())
			}
		} else {
			nama := r.PostFormValue("nama")
			mataKuliah := r.PostFormValue("mataKuliah")
			nilai, _ := strconv.Atoi("nilai")
			getID := r.PostFormValue("id")
			id, _ := strconv.Atoi(getID)
			nilaiMahasiswa = NilaiMahasiswa{
				Nama:       nama,
				MataKuliah: mataKuliah,
				Nilai:      uint(nilai),
				ID:         uint(id),
			}
		}
		nilaiMahasiswa.IndeksNilai = hitungIndeks(nilaiMahasiswa.Nilai)

		semuaDataMahasiswa = append(semuaDataMahasiswa, nilaiMahasiswa)
		dataNilai, _ := json.Marshal(nilaiMahasiswa)
		w.Write(dataNilai)
		return
	}
	http.Error(w, "NOT FOUND", http.StatusNotFound)
}

func hitungIndeks(nilai uint) string {
	switch {
	case nilai >= 80:
		return "A"
	case nilai >= 70 && nilai < 80:
		return "B"
	case nilai >= 60 && nilai < 70:
		return "C"
	case nilai >= 50 && nilai < 60:
		return "D"
	default:
		return "E"
	}
}

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		uname, pwd, ok := r.BasicAuth()
		if !ok {
			w.Write([]byte("Username atau Password tidak boleh kosong"))
			return
		}

		if uname == "admin" && pwd == "admin" {
			next.ServeHTTP(w, r)
			return
		}
	})
}

// soal 2
func tampilNilai(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		dataNilai, err := json.Marshal(semuaDataMahasiswa)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(dataNilai)
		return
	}
	http.Error(w, "Error", http.StatusNotFound)
}
