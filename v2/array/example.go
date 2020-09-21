package main

import "fmt"

func main() {
	var arr [5]int

	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	arr[3] = 4
	arr[4] = 5

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	fmt.Println(arr[2])

	arrStr := [2]string{"apple", "banana"}
	fmt.Println(arrStr)
	fmt.Println(arrStr[1])

	arrMulti := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(arrMulti[0][2])

	genMap := [][]int{} // Make the outer slice with size 0
	for i := 0; i < 2; i++ {
		m := []int{} // Make one inner slice per iteration with size 0
		for j := 0; j < 2; j++ {
			m = append(m, 1) // Append to the inner slice
		}
		genMap = append(genMap, m) // Append to the outer slice
	}
}
