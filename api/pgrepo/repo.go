package pgrepo

import (
	"log"

	"github.com/bromivipo/marketplace/api/definitions"
	sqlqueries "github.com/bromivipo/marketplace/api/pgrepo/sql_queries"
	"github.com/jackc/pgx"
)


func GetConnection() *pgx.Conn {
	config := pgx.ConnConfig{Host: "localhost", Port: 5432, Database: "marketplace", User: "misha", Password: "1111"}
	conn, err := pgx.Connect(config)
	if err != nil {
		log.Println("ERROR: Cannot establish connection")
		panic(err)
	}
	return conn
}

func GetProductById(id int) (*definitions.ProductItem) {
	conn := GetConnection()
	row := conn.QueryRow(sqlqueries.SELECT_PRODUCT_BY_ID, id)
	product := definitions.ProductItem{}
	if err := row.Scan(&product.ID, &product.Name, &product.Price, &product.Category); err != nil {
		log.Printf("ERROR: %v", err)
		return nil
	}
	return &product

}

func GetProducts() ([]definitions.ProductItem) {
	conn := GetConnection()
	rows, err := conn.Query(sqlqueries.SELECT_PRODUCTS)
	resp := []definitions.ProductItem{}

	if err != nil {
		log.Printf("Error in GetProducts: %v", err)
        return resp
    }

	for rows.Next() {
		product := definitions.ProductItem{}
		if err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.Category); err != nil {
			log.Printf("ERROR: %v", err)
			return resp
		}
		resp = append(resp, product)
	}
	return resp

}
