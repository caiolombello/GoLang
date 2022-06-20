package main

import (
	"fmt"
)

type Product struct {
	Name          string
	Description   string
	PurchasePrice float64
}

func (p Product) getSalesPrice() float64 {
	return p.PurchasePrice * 2
}

func main() {
	product1 := Product{
		Name:          "Laptop",
		Description:   "Laptop description",
		PurchasePrice: 1000.00,
	}
	fmt.Println(product1.getSalesPrice())
}
