package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Mahasiswa struct {
	NoBP				string `json:"noBp"`
	Nama        string `json:"nama"`
	Alamat      struct {
		Jalan     string `json:"jalan"`
		Kecamatan string `json:"kecamatan"`
		Kabupaten string `json:"kab-kota"`
		Provinsi  string `json:"provinsi"`
	} `json:"alamat"`
	Jurusan     string `json:"jurusan"`
	Prodi				string `json:"prodi"`
	Email				string `json:"email"`
	Nilai []Nilai `json:"Nilai"`
}

type Nilai struct {
	BpMhs  			string `json:"bp_mahasiswa"`
	IdMatkul 		string `json:"id_matkul"`
	NamaMK			string `json:"nama_mk"`
	NipDosen		string `json:"nip_dosen"`
	NamaDosen		string `json:"nm_dosen"`
	Nilai       float32 `json:"nilai"`
	Semester    int8    `json:"semester"`
}

func main() {
	url := "http://localhost:8080/scores/1811082018"
	spaceClient := http.Client{
		Timeout: time.Second * 2, // Timeout after 2 seconds
	}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("User-Agent", "spacecount-tutorial")
	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}
	if res.Body != nil {
		defer res.Body.Close()
	}
	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	mahasiswa := Mahasiswa{}
	jsonErr := json.Unmarshal(body, &mahasiswa)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	fmt.Println(mahasiswa.NoBP)
	fmt.Println(mahasiswa.Nama)
	fmt.Println(mahasiswa.Jurusan)
	fmt.Println(mahasiswa.Prodi)
	fmt.Println(mahasiswa.Email)
	for _, nilai := range mahasiswa.Nilai {
		fmt.Println("No BP ", nilai.BpMhs)
		fmt.Println("Id Matkul ", nilai.IdMatkul)
		fmt.Println("Matkul ", nilai.NamaMK)
		fmt.Println("Nip Dosen ", nilai.NipDosen)
		fmt.Println("Nilai ", nilai.Nilai)
		fmt.Println("Semester ", nilai.Semester)
	}	
}