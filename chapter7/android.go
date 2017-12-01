package main

import "fmt"

type Person struct {
	Name string
}

type Android struct {
	Person //embeded type or anonymous field
	Model string
}



func (p *Person) Talk() {
	fmt.Println("Hello, my name is " + p.Name)
}

func main() {
	a := Android{Person{"Elmo"}, "Human"}
	
	a.Talk()
	// or
	a.Person.Talk()
}