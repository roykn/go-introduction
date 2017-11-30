package main

import "fmt"

func main() {
	nextEven := makeEvenGenerator()

	fmt.Println(nextEven())
	fmt.Println(nextEven())
	fmt.Println(nextEven())
}

// returns a function that generates even numbers
// i (unlike normal local variables) persists between calls
func makeEvenGenerator() func() uint {
	i := uint(0)

	return func() uint {
		i += 2
		return i
	}
}