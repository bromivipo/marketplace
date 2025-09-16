package pgrepo

import (
	"log"

	"github.com/bromivipo/marketplace/partner_api/pgrepo/sqlqueries"
	"github.com/jackc/pgx"
)


func GetConnection() *pgx.Conn {
	config := pgx.ConnConfig{Host: "localhost", Port: 5432, Database: "partners", User: "misha", Password: "1111"}
	conn, err := pgx.Connect(config)
	if err != nil {
		log.Println("ERROR: Cannot establish connection")
		panic(err)
	}
	return conn
}

func RegisterPartner(name string, token string) error {
	conn := GetConnection()
	_, err := conn.Exec(sqlqueries.INSERT_NEW_PARTNER, name, token)
	return err
}

func SelectPartner(token string) (*int) {
	conn := GetConnection()
	var id int 
	row := conn.QueryRow(sqlqueries.SELECT_PARTNER_BY_TOKEN, token)
	if err := row.Scan(&id); err != nil {
		log.Printf("ERROR: %v", err)
		return nil
	}
	return &id
}