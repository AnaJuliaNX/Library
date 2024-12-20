package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func OpenDB() (*sql.DB, error) {
	DNS := "root:projetosenha@tcp(localhost:33066)/biblioteca?charset=utf8&parseTime=True&loc=Local"

	db, err := sql.Open("mysql", DNS)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	return db, nil
}
