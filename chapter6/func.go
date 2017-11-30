package main

import "fmt"

func main() {
	x := []float64 { 98, 93, 77, 82, 83 }
	fmt.Println(average(x))
}

func average(xs []float64) float64 {
	total := .0
	
	for _, value := range xs {
		total += value // value is like x[i]
	}

	return total / float64(len(xs))
}
