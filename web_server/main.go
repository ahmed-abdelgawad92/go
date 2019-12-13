package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type User struct {
	Nr      int
	Vorname string
}

var db *sql.DB

// init is invoked before main()
func init() {
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
	db, err := sql.Open("mysql", DNS)

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	values := r.FormValue("id")
	fmt.Println(values)
	// w.Write([]byte(`{ message: "Success" }`))
}

func main() {
	println(db)
	users, err := db.Query("SELECT nr, vorname FROM user WHERE abrechnung = 'onlinebuchung'")

	if err != nil {
		panic(err.Error())
	}

	println(users)
	// for users.Next() {
	// 	var user User
	// 	err := users.Scan(&user.Nr, &user.Vorname)
	// 	if err != nil {
	// 		panic(err.Error())
	// 	}
	// 	fmt.Println(user.Vorname)
	// }
	mux := http.NewServeMux()

	mux.HandleFunc("/", indexHandler)
	log.Fatal(http.ListenAndServe(":15234", mux))
}
