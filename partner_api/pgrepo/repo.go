package pgrepo

import (
	"log"
	"os"
	"strconv"

	"github.com/bromivipo/marketplace/partner_api/pgrepo/sqlqueries"
	"github.com/jackc/pgx"
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
	config := pgx.ConnConfig{Host: GetEnvOrDefault("DB_HOST", "localhost"), Port: uint16(port), Database: GetEnvOrDefault("DB_NAME", "partners"), User: GetEnvOrDefault("DB_USER", "misha"), Password: GetEnvOrDefault("DB_PASSWORD", "1111")}
	conn, err := pgx.Connect(config)
	log.Println("CONFIG", config)
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