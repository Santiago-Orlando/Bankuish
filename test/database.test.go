package test

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func GetConnection() (*sql.DB) {

	host := "host" + "=" + "localhost" + " "
	port := "port" + "=" + "3002" + " "
	user := "user" + "=" + "postgres" + " "
	password := "password" + "=" + "your_password" + " "
	name := "dbname" + "=" + "test" + " "
	
	connStr := host + port + user + password + name + "sslmode=disable" 
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	return db
}