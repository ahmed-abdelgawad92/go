package config

import (
	"database/sql"
	"fmt"
		_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"os"
)

//GetDB return a DB instance
func GetDB() (db *sql.DB, err error) {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
	//retreive the connection values
	username := os.Getenv("MYSQL_LOGIN_USER")
	password := os.Getenv("MYSQL_LOGIN_PASSWORD")
	dbname := os.Getenv("MYSQL_LOGIN_DB")
	ip := os.Getenv("MYSQL_LOGIN_OUTER_IP")
	port := os.Getenv("MYSQL_LOGIN_OUTER_PORT")

	//creating the DNS
	DNS := fmt.Sprintf(`%s:%s@tcp(%s:%s)/%s`, username, password, ip, port, dbname)
	return sql.Open("mysql", DNS)
}
