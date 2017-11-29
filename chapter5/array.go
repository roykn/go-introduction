package main

import "fmt"

func main() {
	x := [5]float64 { 98, 93, 77, 82, 83 }

	total := .0

	// 1. Iteration
	// for i := 0; i < len(x); i++ {
	// 	total += x[i]
	// }

	// 2. Iteration
	// for i, value := range x {
	// 	total += value // value is like x[i]
	// }

	// 3. Iteration
	// compiler will complain unused variable i, 
	// so we tell the compiler we don't need it by writing _
	for _, value := range x {
		total += value // value is like x[i]
	}

	fmt.Println(total / float64(len(x)))
}