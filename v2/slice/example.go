package main

import "fmt"

func main() {
	mySlice := []int{1,2,3,4,5}

	for i, v := range mySlice {
		fmt.Println(i, v)
	}

	mySliceStr := []string{"Ali", "Depa", "Alfi", "Jojo", "Aneki"}

	mySliceStr = append(mySliceStr, "Arya")

	for _, v := range mySliceStr{
		fmt.Println(v)
	}
}
