package main

import "fmt"

func main() {
	var mapPerson = exampleMap()
	for k, v := range mapPerson {
		fmt.Println(k, v)
	}
	alfi, ok := mapPerson[3]

	if ok {
		fmt.Println(alfi)
	} else {
		fmt.Println("alfi tidak ada")
	}

	for _, v := range mapStruct() {
		fmt.Println(v.Name)
	}
}

func exampleMap() map[int]string {
	var mapPerson map[int]string
	mapPerson = make(map[int]string)

	mapPerson[1] = "Ali"
	mapPerson[2] = "Depa"

	return mapPerson
}

func mapStruct() map[int]Player {
	mapPlayer := make(map[int]Player)

	mapPlayer[1] = Player{Id: 1, Name: "Alfi"}
	mapPlayer[2] = Player{Id: 2, Name: "Jojo"}

	return mapPlayer
}

type Player struct {
	Id   int
	Name string
}
