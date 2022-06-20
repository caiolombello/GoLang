package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"net/http"
)

type Product struct {
	Name  string
	Price float64
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":9191", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	product := Product{Name: "IPhone", Price: 1000}
	persistProduct(product)
	w.Write([]byte("Hello, world!"))
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
