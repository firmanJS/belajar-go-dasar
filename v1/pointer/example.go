package main

import "fmt"

func changePoint(point *int) (hasil int) {
	*point = 200
	return 
}

func main() {
	var point = 100

	changePoint(&point)

	fmt.Println(point)
}
