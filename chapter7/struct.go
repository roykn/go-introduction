package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x float64
	y float64
	r float64
}

// or
// type Circle struct {
// 	x, y, r float64
// }

func main() {
	// create a new instance 
	// var c Circle
	// or
	// c := new(Circle) // return a pointer
	// or
	c := Circle{ 0,0,5 }
	fmt.Println(circleArea(&c))
	fmt.Println(c.area())
} 

func circleArea(c *Circle) float64 {
	return math.Pi * c.r * c.r
}

func (c *Circle) area() float64 { // receiver
	return math.Pi * c.r * c.r
}