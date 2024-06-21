package main

import "fmt"

type Person struct{}

func main() {
	// slices
	s := []int{1}
	fmt.Println()
	fmt.Printf("initial slice first value: %d\n", s[0])

	sliceAlter(s)
	fmt.Printf("altered non-pointer first value: %d\n", s[0])
	// functions
	// maps
	m := map[string]int{"a": 1}
	fmt.Printf("initial map value of key 'a' %d \n", m["a"])
	mapAlter(m)
	fmt.Printf("altered non-pointer map value of 'a' %d \n", m["a"])
	// channels
}

func sliceAlter(s []int) {
	s[0] = 100
}

func mapAlter(m map[string]int) {
	m["a"] = 100
}
