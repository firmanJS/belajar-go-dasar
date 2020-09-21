package main

import "fmt"

func main() {
	p := &Player{
		Id:   1,
		Name: "Depa",
		Age:  20,
	}

	fmt.Println(p.getName())
	p.changeName("ALFI")
	fmt.Println(p.getName())
}

type Player struct {
	Id   int
	Name string
	Age  int
}

func (p *Player) changeName(newName string) {
	p.Name = newName
}

func (p *Player) getId() int {
	return p.Id
}

func (p *Player) getName() string {
	return p.Name
}

func (p *Player) getAge() int {
	return p.Age
}
