package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("JSON in GO")
	palyer1 := Player{ID: 1, Name: "Alfi"}

	p1, err := json.Marshal(palyer1)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(string(p1))

	fmt.Println("====================")

	data := []byte(`{"id":2,"name":"Depa"}`)

	var p2 Player

	err = json.Unmarshal(data, &p2)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(p2)

}

type Player struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
