package main

import "fmt"


// simple loop
// func main() {
// 	i := 0
// 	for i < 5 {
// 		fmt.Println("Nilai i ",i)
// 		i = i + 1
// 	}
// }


//deret angka
func main(){
	total := 0

	for num := 1; num <= 100; num++ {
		total = total + num
	}

	fmt.Println("Total ",total)
}