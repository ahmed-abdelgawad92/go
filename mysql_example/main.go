package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type User struct {
	Nr         int
	Vorname    string
	Abrechnung string
}

// init is invoked before main()
func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {

	//retreive the connection values
	username := os.Getenv("MYSQL_LOGIN_USER")
	password := os.Getenv("MYSQL_LOGIN_PASSWORD")
	dbname := os.Getenv("MYSQL_LOGIN_DB")
	ip := os.Getenv("MYSQL_LOGIN_OUTER_IP")
	port := os.Getenv("MYSQL_LOGIN_OUTER_PORT")

	//creating the DNS
	DNS := fmt.Sprintf(`%s:%s@tcp(%s:%s)/%s`, username, password, ip, port, dbname)
	db, err := sql.Open("mysql", DNS)

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	// users, err := db.Query("SELECT nr, name, abrechnung FROM user WHERE abrechnung = 'onlinebuchung'")

	// if err != nil {
	// 	panic(err.Error())
	// }

	// for users.Next() {
	// 	var user User
	// 	err := users.Scan(&user.nr, &user.vorname, &user.abrechnung)
	// 	if err != nil {
	// 		panic(err.Error())
	// 	}
	// 	fmt.Println(user.VORNAME)
	// }

	//prepared statements
	stmt, err := db.Prepare("SELECT * FROM user WHERE abrechnung = ? LIMIT 10")
	if err != nil {
		panic(err.Error())
	}

	usersPrep, err := stmt.Query("quartal")
	if err != nil {
		panic(err.Error())
	}

	defer stmt.Close()

	var data []User

	for usersPrep.Next() {
		var user User
		err = usersPrep.Scan(&user)
		if err != nil {
			panic(err.Error())
		}
		data = append(data, user)
	}

	jsonData, err := json.Marshal(data)

	fmt.Println(string(jsonData))
}
