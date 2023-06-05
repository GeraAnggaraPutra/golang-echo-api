package db

import (
	"database/sql"
	"fmt"
	"go-echo/config"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

func Init() *sql.DB {
	conf := config.GetConfig()

	connectionString := conf.DB_USERNAME + ":" + conf.DB_PASSWORD + "@tcp(" + conf.DB_HOST + ":" + conf.DB_PORT + ")/" + conf.DB_NAME + "?parseTime=true"

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database connection success")
	return db
}
