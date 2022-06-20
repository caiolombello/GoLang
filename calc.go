package main

import (
	"fmt"
)

func calculate(x, y float64, op string) float64 {
	if op == "+" {
		return x + y
	} else if op == "-" {
		return x - y
	} else if op == "/" {
		return x / y
	} else if op == "*" || op == "x" {
		return x * y
	} else {
		fmt.Println("Operation don't found")
		return 0
	}
}

func main() {
	var x, y float64
	var op string
	fmt.Println("Calculate two numbers:")
	fmt.Scan(&x, &op, &y)
	fmt.Println(calculate(x, y, op))
}
