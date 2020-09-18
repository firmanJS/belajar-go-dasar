package main

import "fmt"

func cariWarna(warna string) (hasil string, boolean bool) {

	if warna == "hijau" {
		hasil = "betul yang dicari"
		boolean = true
	} else if warna == "biru" {
		hasil = "kemungkinan lain"
		boolean = true
	} else {
		hasil = "bukan yang dicari"
		boolean = false
	}

	return
}

func main(){
	hasil, boolean := cariWarna("merah")
	fmt.Println(hasil, boolean)
}