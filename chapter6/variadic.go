package main

import "fmt"

func main() {
	x := []int{1,2,3,4,5,6,7,8,9}
	fmt.Println(add(x...))
}

// variadic parameters
// zero or more parameters of type int
func add(values ...int) int {
	total := 0
	for _, v := range values {
		total += v
	}
	return total
}