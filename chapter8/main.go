package main

import (
	"chapter8/math" // add to GOPATH
	"fmt"
)

func main() {
	avrg := math.Average([]float64{1, 2, 3})
	fmt.Println(avrg)
}
