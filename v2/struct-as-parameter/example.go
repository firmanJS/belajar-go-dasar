package main

import "fmt"

type Vector struct {
	x int
	y int
}

type Player struct {
	ID   int
	Name string
}

func main() {
	var v Vector

	v.x = 1
	v.y = 10

	fmt.Println(v)
	fmt.Println("X =", v.x)
	fmt.Println("Y =", v.y)

	player1 := Player{ID: 1, Name: "Depa"}
	fmt.Println(player1.ID)
	fmt.Println(player1.Name)

}
