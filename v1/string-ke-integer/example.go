package main

import (
	"fmt"
	"strconv"
)

func main() {
	gaji := "20000"
	gaji2,_:= strconv.Atoi(gaji)
	bonus := gaji2 +  1000
	fmt.Println("bonus", bonus)
}
