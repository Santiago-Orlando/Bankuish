package config

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

func GetConnection() (*sql.DB) {

	host := "host" + "=" + os.Getenv("HOST") + " "
	port := "port" + "=" + os.Getenv("DB_PORT") + " "
	user := "user" + "=" + os.Getenv("USER") + " "
	password := "password" + "=" + os.Getenv("PASSWORD") + " "
	name := "dbname" + "=" + os.Getenv("DB_NAME") + " "
	
	connStr := host + port + user + password + name + "sslmode=disable" 
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	return db
}