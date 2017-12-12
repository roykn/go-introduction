package main

import (
	"fmt"
	"intro_go/chapter8/math" // add to GOPATH
)

func main() {
	avrg := math.Average([]float64{1, 2, 3})
	fmt.Println(avrg)
}
