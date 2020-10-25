package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"io/ioutil"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"gopkg.in/yaml.v2"
)

var db *sql.DB
var err error

type YamlConfig struct {
	Connection struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Database string `yaml:"database"`
	}
}

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

func getScores(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var mahasiswa Mahasiswa
	var nilai Nilai
	params := mux.Vars(r)

	sql := `SELECT
				nobp,
				nama,
				IFNULL(jalan,'') jalan,
				IFNULL(kecamatan,'') kecamatan,
				IFNULL(kabupaten,'') kabupaten,
				IFNULL(provinsi,'') provinsi,
				IFNULL(jurusan,'') jurusan,
				IFNULL(prodi,'') prodi,
				IFNULL(email,'') email				
			FROM mahasiswa WHERE nobp IN (?)`
	result, err := db.Query(sql, params["id"])
	defer result.Close()

	if err != nil {
		panic(err.Error())
	}
	for result.Next() {
		err := result.Scan(&mahasiswa.NoBP, &mahasiswa.Nama, &mahasiswa.Alamat.Jalan, &mahasiswa.Alamat.Kecamatan, &mahasiswa.Alamat.Kabupaten, &mahasiswa.Alamat.Provinsi, &mahasiswa.Jurusan, &mahasiswa.Prodi, &mahasiswa.Email)
		if err != nil {
			panic(err.Error())
		}
		sqlDetial := `SELECT
				nilai.bp_mhs,
				nilai.id_matkul,
				matkul.nama_mk,
				nilai.nip_dosen,
				dosen.nama,
				nilai.nilai,
				nilai.semester
				FROM
					nilai
					INNER JOIN mahasiswa
						ON (nilai.bp_mhs = mahasiswa.nobp)
						INNER JOIN dosen
							ON (nilai.nip_dosen = dosen.nip)
							INNER JOIN matkul
								ON (nilai.id_matkul = matkul.id)
				WHERE nilai.bp_mhs	= ?`
		bpMhs := &mahasiswa.NoBP
		fmt.Println(*bpMhs)
		resultDetail, errDet := db.Query(sqlDetial, *bpMhs)
		defer resultDetail.Close()
		if errDet != nil {
			panic(err.Error())
		}
		for resultDetail.Next() {	
			err := resultDetail.Scan(&nilai.BpMhs, &nilai.IdMatkul, &nilai.NamaMK, &nilai.NipDosen, &nilai.NamaDosen, &nilai.Nilai, &nilai.Semester)
			if err != nil {
				panic(err.Error())
			}
			mahasiswa.Nilai = append(mahasiswa.Nilai, nilai)	
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mahasiswa)
}


// Main function
func main() {
	yamlFile, err := ioutil.ReadFile("../Yaml/config.yml")
	if err != nil {
		fmt.Printf("Error reading YAML file: %s\n", err)
		return
	}
	var yamlConfig YamlConfig
	err = yaml.Unmarshal(yamlFile, &yamlConfig)
	if err != nil {
		fmt.Printf("Error parsing YAML file: %s\n", err)
	}

	host := yamlConfig.Connection.Host
	port := yamlConfig.Connection.Port
	user := yamlConfig.Connection.User
	pass := yamlConfig.Connection.Password
	dabes := yamlConfig.Connection.Database

	var (
		mySQL = fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", user, pass, host, port, dabes)
	)

	db, err = sql.Open("mysql", mySQL)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	// Init router
	r := mux.NewRouter()

	// Route handles & endpoints
	r.HandleFunc("/scores/{id}", getScores).Methods("GET")
	
	fmt.Println("Server on :8080")
	// Start server
	log.Fatal(http.ListenAndServe(":8080", r))
	
}
