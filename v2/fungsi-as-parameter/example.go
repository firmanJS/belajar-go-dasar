package main

import "fmt"

func main() {
	f := func(v string) bool {
		return v == "Golang"
	}

	result := match("Golang", f)

	fmt.Println(result)
}

func match(v string, f func(string) bool) bool {

	if f(v) {
		return true
	}
	return false
}