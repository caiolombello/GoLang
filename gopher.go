package main

import (
	"fmt"
)

// Scope: Package Level
var y = "Hello, World!"

func main() {

	// Scope: Code Block
	// Declaração
	x := 10

	// Tipagem automática
	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("y: %v, %T\n", y, y)

	x, z := 20, 30
	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("z: %v, %T\n", z, z)

	// Atribuição
	z = 10
	fmt.Printf("z: %v, %T\n", z, z)
	
}
