package main

import (
	"fmt"
	"os"
	"strconv"
)

type Persons struct {
	Nama, Alamat, Pekerjaan, Alasan string
}

func datas() map[int]Persons {
	person := make(map[int]Persons, 2)
	person[1] = Persons{
		Nama:      "Arwin Thoriq R",
		Alamat:    "Blora",
		Pekerjaan: "Software Engineer",
		Alasan:    "Karena Mudah untuk dipelajari",
	}
	person[2] = Persons{
		Nama:      "Karta Kusuma",
		Alamat:    "Jakarta",
		Pekerjaan: "Software Engineer",
		Alasan:    "Karena Mudah untuk dipelajari",
	}
	person[3] = Persons{
		Nama:      "Yunus Fajri",
		Alamat:    "Bogor",
		Pekerjaan: "Software Engineer",
		Alasan:    "Karena Mudah untuk dipelajari",
	}
	person[4] = Persons{
		Nama:      "Ansharullah Widiansyah",
		Alamat:    "Bekasi",
		Pekerjaan: "Software Engineer",
		Alasan:    "Karena Mudah untuk dipelajari",
	}

	return person
}

func getPersonByKey(key int, person map[int]Persons) {
	fmt.Printf("Nama: %s\n", person[key].Nama)
	fmt.Printf("Alamat: %s\n", person[key].Alamat)
	fmt.Printf("Pekerjaan: %s\n", person[key].Pekerjaan)
	fmt.Printf("Alasan Memilih Golang: %s\n", person[key].Alasan)
}

func tugasempat() {

	if len(os.Args) != 2 {
		fmt.Println("masukkan go run . 1(data absen")
		os.Exit(1)
	}

	number, err := strconv.Atoi(os.Args[1])
	if number <= 0 || err != nil {
		fmt.Println("Input berupa angka")
		os.Exit(2)
	}

	datas := datas()

	if number > 0 && number <= len(datas) {
		getPersonByKey(number, datas)
	} else {
		fmt.Println("absen tidak tersedia")
	}
}
