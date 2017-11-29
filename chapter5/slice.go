package main

import "fmt"

func main() {
	
	// different ways to create slices
	// x := []float64
	// or
	// x := make([]float64, 5, 10)
	// or
	arr := [5]float64 { 1,2,3,4,5 }
	slc1 := arr[0:5]

	// append -> adds elements to the end of a slice
	slc2 := append(slc1, 6, 7, 8, 9)
	fmt.Println(slc1, slc2)

	// copy
	slc3 := []float64 { 10, 11, 12 }
	copy(slc2, slc3) // (dst, src)
	fmt.Println(slc2, slc3)
}