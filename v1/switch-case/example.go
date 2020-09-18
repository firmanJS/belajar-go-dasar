package main

import "fmt"

func cariWarna(warna string) (hasil string) {

	switch {
	case warna == "hijau":
		hasil = "betul yang dicari"
	case warna == "merah":
		hasil = "betul yang dicari"
	case warna == "biru":
		hasil = "betul yang dicari"
	default:
		hasil = "bukan yang dicari"
	}

	return
}

func main() {
	hasil := cariWarna("semlehoy")
	fmt.Println(hasil)
}
