package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/souvikhaldar/noteledger/pkg/config"
)

type DB struct {
	mysql *sql.DB
}

func NewDB() *DB {
	config := config.GetConfig("/Users/souvik/Development/go/src/github.com/souvikhaldar/noteledger/conf.json")
	log.Println("Username: ", config.DB_Username)
	db, err := sql.Open("mysql", config.DB_Username+":"+config.DB_Password+"@tcp(127.0.0.1:3306)/"+config.DB_Table_Name)
	if err != nil {
		log.Fatal("Unable to connect to db: ", err)
	}
	return &DB{
		mysql: db,
	}
}
