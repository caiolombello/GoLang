package main

import (
	"fmt"
)

func soma(a int, b int) int {
	return a + b
}

func main() {
	var n1, n2 int

	fmt.Println("Digite dois números para soma:")
	fmt.Scan(&n1, &n2)
	fmt.Println(soma(n1, n2))
}