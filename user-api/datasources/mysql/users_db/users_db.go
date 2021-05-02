package users_db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	Client *sql.DB
)

func init() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", "root", "keep1234", "localhost:3306", "users_db")
	var err error
	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil {

		panic(err)
	}
	if err = Client.Ping(); err != nil {

		panic(err)
	}

	log.Println("database successfully configured")
}
