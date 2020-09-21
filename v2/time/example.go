package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Time in GO")

	t := time.Now()

	fmt.Println(t)
	fmt.Println(t.Year())
	fmt.Println(t.Month())
	fmt.Println(t.Day())

	timeString := "January 21, 2020"
	from := "January 2, 2006"

	resTime, err := time.Parse(from, timeString)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resTime)
	fmt.Println(resTime.Year())

	t1 := time.Date(2015, 2, 1, 12, 0, 0, 0, time.UTC)
	t2 := time.Date(2015, 2, 1, 12, 0, 0, 0, time.UTC)

	eq := t1.Equal(t2)

	fmt.Println(t1)
	fmt.Println(eq)

	layout := "2020-01-02"

	resString := t.Format(layout)

	fmt.Println(resString)
}
