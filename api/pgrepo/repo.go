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


func GetProductById(id int) (*generated.ProductItem) {
	conn := GetConnection()
	row := conn.QueryRow(sqlqueries.SELECT_PRODUCT_BY_ID, id)
	product := generated.ProductItem{}
	if err := row.Scan(&product.Id, &product.Name, &product.Price, &product.LeftInStock, &product.ProviderId, &product.Category); err != nil {
		log.Printf("ERROR: %v", err)
		return nil
	}
	return &product

}

func GetProducts() ([]generated.ProductItem) {
	conn := GetConnection()
	rows, err := conn.Query(sqlqueries.SELECT_PRODUCTS)
	resp := []generated.ProductItem{}

	if err != nil {
		log.Printf("Error in GetProducts: %v", err)
        panic(err)
    }

	for rows.Next() {
		product := generated.ProductItem{}
		if err := rows.Scan(&product.Id, &product.Name, &product.Price, &product.LeftInStock, &product.ProviderId, &product.Category); err != nil {
			log.Printf("ERROR: %v", err)
			return resp			
		}
		resp = append(resp, product)
	}
	return resp

}


func RegisterUser(username string, password string) (error) {
	conn := GetConnection()
	_, err := conn.Exec(sqlqueries.INSERT_NEW_USER, username, password)
	return err
}


func GetUserPassword(username string) (*string) {
	conn := GetConnection()
	row := conn.QueryRow(sqlqueries.SELECT_USER, username)
	var password string
	if err := row.Scan(&password); err != nil {
		log.Printf("ERROR: %v", err)
		return nil
	}
	return &password
}