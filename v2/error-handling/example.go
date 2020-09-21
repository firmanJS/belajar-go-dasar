package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {
	iStr := "a"

	i, err := strconv.Atoi(iStr)

	if err != nil {
		fmt.Println("Terjadi Kesalahan : ", err.Error())
	} else {
		fmt.Println(i)
	}

	r, err := Div(10, 2)
	if err != nil {
		fmt.Println("Terjadi Kesalahan : ", err.Error())
	} else {
		fmt.Println(r)
	}

	//inline error handlig
	if x, err := Div(10, 5); err != nil {
		fmt.Println("Terjadi Kesalahan : ", err.Error())
	} else {
		fmt.Println(x)
	}
}

func Div(x, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("Tidak bisa membagi dengan 0")
	} else {
		return x / y, nil
	}
}
