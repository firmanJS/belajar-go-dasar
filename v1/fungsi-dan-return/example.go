package main

import (
	"fmt"
	"strconv"
)

var text = "pesan rahasia ..."

func kali(angka1 int, angka2 int) int {
	return angka1 * angka2
}

func infoMultiReturn(umur int, nama string, status string) (string, string) {
	umurSekarang := strconv.Itoa(umur)

	return "nama : " + nama + " umur :" , umurSekarang + " hubungan : " + status
}

func infoSingleReturn(umur int, nama string, status string) string {
	umurSekarang := strconv.Itoa(umur)

	return "nama : " + nama + " umur :" + umurSekarang + " hubungan : " + status
}

// memberi nama nilai yang di return
func infoReturn(umur int, nama string, status string) (bio string, hubungan string) {
	hubungan = status
	bio = "nama : " + nama

	return
}

//fungsi tanpa return
func noReturn() {
	fmt.Println(text)
}

func main() {
	// infoNama, infoPersonal := infoReturn(12, "depa", "menikah")
	// fmt.Println("hasil perkalian", kali(2, 2))
	// fmt.Println("halo perkenalkan", infoSingleReturn(12, "depa", "menikah"))
	// fmt.Println("halo perkenalkan", infoNama, infoPersonal)
	// fmt.Println("halo perkenalkan", infoNama, infoPersonal)
	// fmt.Println("halo perkenalkan", infoNama, infoPersonal)
	noReturn()
}
