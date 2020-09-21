package main

import "fmt"

var hello = func(name string) string {
	return name
}

func main(){
	fmt.Println(hello("adam"))
}