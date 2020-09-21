package main

import "fmt"

func main() {


	//sample1
	// sample := [2][3]int{{1, 2, 3}, {4, 5, 6}}

	// fmt.Printf("Number of rows in array: %d\n", len(sample))
	// fmt.Printf("Number of columns in array: %d\n", len(sample[0]))
	// fmt.Printf("Total number of elements in array: %d\n", len(sample)*len(sample[0]))

	// fmt.Println("Traversing Array")
	// for _, row := range sample {
	// 	for _, val := range row {
	// 		fmt.Println(val)
	// 	}
	// }

	// sample 2
	// sample := [2][2][3]int{{{1, 2, 3}, {4, 5, 6}}, {{7, 8, 9}, {10, 11, 12}}}

    // fmt.Printf("Length of first dimension: %d\n", len(sample))
    // fmt.Printf("Length of second dimension: %d\n", len(sample[0]))
    // fmt.Printf("Length of third dimension: %d\n", len(sample[0][0]))

    // fmt.Printf("Overall Dimension of the array: %d*%d*%d\n", len(sample), len(sample[0]), len(sample[0][0]))
    // fmt.Printf("Total number of elements in array: %d\n", len(sample)*len(sample[0])*len(sample[0][0]))

    // for _, first := range sample {
    //     for _, second := range first {
    //         for _, value := range second {
    //             fmt.Println(value)
    //         }
    //     }
	// }
	
	//sample 3  Accessing elements of a multi dimensional array
	sample := [2][3]int{{1, 2, 3}, {4, 5, 6}}
    //Print array element
    fmt.Println(sample[0][0])
    fmt.Println(sample[0][1])
    fmt.Println(sample[0][2])
    fmt.Println(sample[1][0])
    fmt.Println(sample[1][1])
    fmt.Println(sample[1][2])
    
    //Assign new values
    sample[0][0] = 6
    sample[0][1] = 5
    sample[0][2] = 4
    sample[1][0] = 3
    sample[1][1] = 2
    sample[1][2] = 1

    fmt.Println()
    fmt.Println(sample[0][0])
    fmt.Println(sample[0][1])
    fmt.Println(sample[0][2])
    fmt.Println(sample[1][0])
    fmt.Println(sample[1][1])
    fmt.Println(sample[1][2])
}
