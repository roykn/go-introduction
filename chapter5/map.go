package main

import "fmt"

func main() {
	m := make(map[string]int)  // [key_type]value_type

	m["key_1"] = 10
	m["key_2"] = 11

	delete(m, "key_1")

	if key, ok := m["key_1"]; ok {
		fmt.Println(key, ok)
	}
}