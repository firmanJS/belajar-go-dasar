package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Environment in GO")

	os.Setenv("PG_HOST", "localhost")

	postgres := os.Getenv("PG_HOST")

	fmt.Println("Postgres Config : ", postgres)
}
