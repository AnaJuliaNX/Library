package repository

import (
	"errors"
	dbconection "library/domains/users/dbConection"
)

func getTotalUsers(name string) (int, error) {
	db, err := dbconection.ConnectDB()
	if err != nil {
		return 0, errors.New("Não foi possível conectar ao banco de dados")
	}
	defer db.Close()

	var query string
	var args []interface{}
	var total int

	if name == "" {
		query = "SELECT COUNT(*) FROM Users"
	} else {
		query = "SELECT COUNT(*) FROM Users WHERE Name LIKE ?"
		args = append(args, "%"+name+"%")
	}

	err = db.QueryRow(query, args...).Scan(&total)
	if err != nil {
		return 0, errors.New("Não foi possível buscar o total de usuários")
	}

	return total, nil
}
