package main

import "fmt"

func main() {
	x := 10

	add := func(values ...int) int {
		total := 0
		for _, v := range values {
			total += v
		}
		return total + x
	}

	fmt.Println(add(1,2,3))
}
