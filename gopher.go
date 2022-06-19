package main

import (
	"fmt"
)

func main() {

	// Declaração
	x := 10
	y := "Hello, World!"

	// Tipagem automática
	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("y: %v, %T\n", y, y)

	// Atribuição
	x = 20
	fmt.Printf("x: %v, %T\n", x, x)
}
