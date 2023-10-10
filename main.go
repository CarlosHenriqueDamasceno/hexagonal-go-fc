package main

import (
	"database/sql"
	"log"

	"github.com/CarlosHenriqueDamasceno/hexagonal-fc/adapters/db"
	"github.com/CarlosHenriqueDamasceno/hexagonal-fc/application"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	connection, _ := sql.Open("sqlite3", "db.sqlite3")
	productDbAdapter := db.NewProductDb(connection)
	productService := application.NewProductService(productDbAdapter)
	product, err := productService.Create("Product Example", 30)
	if err != nil {
		log.Fatal(err.Error())
	}
	productService.Enable(product)
}
