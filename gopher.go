package main

import (
	"fmt"
)

// Package Level: Scope
var y = "Hello, World!"

func main() {

	// Package Level: Code Block
	// Declaração
	x := 10

	// Tipagem automática
	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("y: %v, %T\n", y, y)

	// Atribuição
	x, z := 20, 30
	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("z: %v, %T\n", z, z)
}
