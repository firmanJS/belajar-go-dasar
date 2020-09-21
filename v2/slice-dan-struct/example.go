package main

import "fmt"

type Player struct {
	Id int
	Name string
}

func main() {
	var players []Player

	players = []Player{Player{Id:1, Name:"Ali"}, Player{Id:3, Name:"Depa"}}

	players = append(players, Player{Id:3, Name:"Subarjo"})

	for _, v := range players{
		fmt.Println(v)
	}
}
