package main

import "fmt"

var bonus = 50
func hitungTotal(gaji int) int {
	return gaji + bonus
}

func main() {
	fmt.Println("bonus awal : ", bonus)
	bonus = 200
	fmt.Println("bonus dan gaji : ", hitungTotal(100))
}
