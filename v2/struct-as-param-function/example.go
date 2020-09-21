package main

import "fmt"

func main() {
	p := Player{
		Id: 1,
		Name: "Depa",
		Age: 20,
	}

	printPlayer(p)
}

type Player struct {
	Id   int
	Name string
	Age int
}

func printPlayer(p Player) {
	fmt.Println(p.Id)
	fmt.Println(p.Name)
	fmt.Println(p.Age)
}
