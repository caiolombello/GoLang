package main

import (
	"database/sql"

	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
)

type Product struct {
	Name  string
	Price float64
}

var products []Product

func generateProducts() {
	p1 := Product{Name: "Product 1", Price: 1.99}
	p2 := Product{Name: "Product 2", Price: 2.99}
	products = append(products, p1, p2)
}

func main() {
	generateProducts()
	e := echo.New()
	e.GET("/product", listProducts)
	e.Logger.Fatal(e.Start(":9191"))
}

func listProducts(c echo.Context) error {
	return c.JSON(200, products)
}

func persistProduct(product Product) {
	db, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	stmt, err := db.Prepare("Insert into products(name, price) values($1, $2)")
	if err != nil {
		panic(err)
	}

	stmt.Exec(product.Name, product.Price)
}
