package main

import (
	"database/sql"

	dbAdapter "github.com/adrianokerber/go-hexagonal/adapters/db"
	"github.com/adrianokerber/go-hexagonal/application"
)

func main() {
	db, _ := sql.Open("sqlite3", "sqlite.db")
	productDbAdapter := dbAdapter.NewProductDb(db)
	productService := application.NewProductService(productDbAdapter)

	product, _ := productService.Create("Product Example", 30)
	productService.Enable(product)
}
