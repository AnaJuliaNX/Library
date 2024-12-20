package repository

import (
	"errors"

	dbconection "library/domains/users/dbConection"
	"library/domains/users/entity"
)

func Index(name string, offset, limit int64) ([]entity.User, int, error) {
	db, err := dbconection.ConnectDB()
	if err != nil {
		return nil, 0, errors.New("Não foi possível conectar ao banco de dados")
	}
	defer db.Close()

	stmt, err := db.Query("SELECT BIN_TO_UUID(ID), Name, Email, Phone, Owner FROM Users WHERE Name LIKE ? LIMIT ? OFFSET ?", "%"+name+"%", limit, offset)
	if err != nil {
		return nil, 0, errors.New("Não foi possível buscar usuários")
	}
	defer stmt.Close()

	var users []entity.User
	for stmt.Next() {
		var user entity.User
		err = stmt.Scan(&user.ID, &user.Name, &user.Email, &user.Phone, &user.Owner)
		if err != nil {
			return nil, 0, errors.New("Não foi possível buscar usuários")
		}
		users = append(users, user)
	}

	// Buscar o total de usuários
	total, err := getTotalUsers(name)
	if err != nil {
		return nil, 0, errors.New("Não foi possível buscar o total de usuários")
	}

	return users, total, nil
}
