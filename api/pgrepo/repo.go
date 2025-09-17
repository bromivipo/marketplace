package pgrepo

import (
	"log"
	"os"
	"strconv"

	"github.com/bromivipo/marketplace/api/definitions"
	sqlqueries "github.com/bromivipo/marketplace/api/pgrepo/sql_queries"
	"github.com/jackc/pgx"
	"github.com/shopspring/decimal"
)

func GetEnvOrDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func GetConnection() *pgx.Conn {
	port, _ := strconv.Atoi(GetEnvOrDefault("DB_PORT", "5432"))
	config := pgx.ConnConfig{Host: GetEnvOrDefault("DB_HOST", "localhost"), Port: uint16(port), Database: GetEnvOrDefault("DB_NAME", "marketplace"), User: GetEnvOrDefault("DB_USER", "misha"), Password: GetEnvOrDefault("DB_PASSWORD", "1111")}
	conn, err := pgx.Connect(config)
	if err != nil {
		log.Println("ERROR: Cannot establish connection")
		panic(err)
	}
	return conn
}

func GetProductById(id int) *generated.ProductItem {
	conn := GetConnection()
	row := conn.QueryRow(sqlqueries.SELECT_PRODUCT_BY_ID, id)
	product := generated.ProductItem{}
	if err := row.Scan(&product.Id, &product.Name, &product.Price, &product.LeftInStock, &product.ProviderId, &product.Category); err != nil {
		log.Printf("ERROR: %v", err)
		return nil
	}
	return &product
}

func GetProducts() []generated.ProductItem {
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

func RegisterUser(username string, password string) error {
	conn := GetConnection()
	_, err := conn.Exec(sqlqueries.INSERT_NEW_USER, username, password)
	return err
}

func GetUserPassword(username string) *string {
	conn := GetConnection()
	row := conn.QueryRow(sqlqueries.SELECT_USER, username)
	var password string
	if err := row.Scan(&password); err != nil {
		log.Printf("ERROR: %v", err)
		return nil
	}
	return &password
}

type ErrorReason int

const (
	OutOfStock ErrorReason = iota
	NotFound
)

type UpdateError struct {
	Id     int
	Reason ErrorReason
}

func UpdateProducts(ids []int) *UpdateError {
	conn := GetConnection()
	trx, err := conn.Begin()
	if err != nil {
		log.Printf("ERROR: %v", err)
	}
	for _, id := range ids {
		row := trx.QueryRow(sqlqueries.UPDATE_PRODUCT, id)
		var left_in_stock int
		if err := row.Scan(&left_in_stock); err != nil {
			log.Printf("ERROR: %v", err)
			trx.Rollback()
			return &UpdateError{Id: id, Reason: NotFound}
		}
		if left_in_stock < 0 {
			trx.Rollback()
			return &UpdateError{Id: id, Reason: OutOfStock}
		}
	}
	trx.Commit()
	return nil
}

func InsertOrder(ids []int, username string) error {
	conn := GetConnection()
	var total_amount decimal.Decimal
	err := conn.QueryRow(sqlqueries.SELECT_TOTAL_AMOUNT, ids).Scan(&total_amount)
	if err != nil {
		log.Printf("ERROR: %v", err)
	}
	_, err = conn.Exec(sqlqueries.INSERT_NEW_ORDER, username, ids, total_amount)
	return err
}

func InsertNewProduct(product *generated.ProductToRegister) (id int32, err error) {
	conn := GetConnection()
	price, _ := decimal.NewFromString(product.Price)
	err = conn.QueryRow(sqlqueries.INSERT_NEW_PRODUCT, product.Name, price, product.LeftInStock, product.ProviderId, product.Category).Scan(&id)
	if err != nil {
		log.Printf("ERROR: %v", err)
	}
	return id, err
}

func UpdateProductsAmount(toAdd *generated.AddProductsRequest) *UpdateError {
	conn := GetConnection()
	trx, err := conn.Begin()
	if err != nil {
		log.Printf("ERROR: %v", err)
	}
	for _, product := range toAdd.Products {
		_, err := trx.Exec(sqlqueries.UPDATE_PRODUCT_AMOUNT, product.Id, product.Amount)
		if err != nil {
			log.Printf("ERROR: %v", err)
			trx.Rollback()
			return &UpdateError{Id: int(product.Id), Reason: NotFound}
		}
	}
	trx.Commit()
	return nil
}
