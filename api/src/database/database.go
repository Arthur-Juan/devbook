package database

import (
	"api/src/config"
	"database/sql"
	_ "github.com/go-sql-driver/mysql" //DRIVER
	"log"
)

func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", config.StrConn)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	if err = db.Ping(); err != nil {
		err := db.Close()
		if err != nil {
			return nil, err
		}
		return nil, err

	}

	return db, nil
}
