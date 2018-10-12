package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Hello World\n")
}

// Add returns the addition of 2 integers
func Add(a, b int) int {
	return a + b
}

// TryStuff lets you know if you passed in true or false
// If false, you'll get an error
func TryStuff(t bool) (bool, error) {
	if t {
		return t, nil
	}

	return t, fmt.Errorf("Error")
}
