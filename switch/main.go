package main

import "fmt"

func main() {
	switch ok := false; {
	case true:
		fmt.Println("true", ok)
	case false:
		fmt.Println("false", ok)
	default:
		fmt.Println("default")
	}
}
