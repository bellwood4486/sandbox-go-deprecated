package main

import "fmt"

func main() {
	var a []int
	b := []int{}

	fmt.Printf("a: %v, b: %v\n", a, b)
	fmt.Printf("len(a): %v, len(b): %v\n", len(a), len(b))
	fmt.Printf("a is nil?: %v\n", a == nil) // -> true
	fmt.Printf("b is nil?: %v\n", b == nil) // -> false
}
