package main

import (
	"fmt"
)

func main() {
	hello := make(chan string, 4)
	hello <- "Hello"
	hello <- "Golang"
	hello <- "Babang"
	hello <- "Tampan"

	close(hello)

	for v := range hello {
		fmt.Println(v)
	}
}
