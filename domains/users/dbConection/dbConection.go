package dbconection

import (
	"database/sql"
	"errors"
	"library/database"
)

func ConnectDB() (*sql.DB, error) {
	db, err := database.OpenDB()
	if err != nil {
		return nil, errors.New("Não foi possivel conectar ao banco de dados")
	}
	return db, nil
}
