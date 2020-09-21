package main

import (
	"fmt"
)

func main() {
	fmt.Println("Defer in Go")

	for i := 1; i <= 5; i++ {
		defer process(i)
	}
}

func process(id int) {
	fmt.Printf("Prosess %d\n", id)
}
