package main

import (
	"flag"
	"fmt"
)

func main() {
	username := flag.String("U", "", "your usernmae")
	password := flag.String("P", "", "your password")

	flag.Parse()

	result := login(*username, *password)

	if result {
		fmt.Println("login Success")
	} else {
		fmt.Println("login failed")
	}
}

func login(username, password string) bool {
	if username == "depa" && password == "depa" {
		return true
	} else {
		return false
	}
}
