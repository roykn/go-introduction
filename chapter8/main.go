package main

import (
	"fmt"
	"go-introduction/chapter8/math"
)

func main() {
	avrg := math.Average([]float64{1, 2, 3})
	fmt.Println(avrg)
}
