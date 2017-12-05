package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Join([]string {"a", "b", "c", "d"}, "-"))
	fmt.Println(strings.Contains("test", "es"))
	fmt.Println(strings.Split("a,b,c,d,e,v,g,k,l,j,g,d,f,g,s,f", ","))
}