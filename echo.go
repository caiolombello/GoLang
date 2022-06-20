package main

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
)

type Product struct {
	Name  string  `json:"nome" validate:"required"`
	Price float64 `json:"preco" validate: required`
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
	e.POST("/product", createProduct)
	e.Logger.Fatal(e.Start(":9191"))
}

func listProducts(c echo.Context) error {
	return c.JSON(200, products)
}

func createProduct(c echo.Context) error {
	product := Product{}
	c.Bind(&product)
	err := persistProduct(product)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}
	return c.JSON(http.StatusCreated, product)
}

func persistProduct(product Product) error {
	db, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		return err
	}
	defer db.Close()
	stmt, err := db.Prepare("Insert into products(name, price) values($1, $2)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(product.Name, product.Price)
	if err != nil {
		return err
	}
	return nil
}
