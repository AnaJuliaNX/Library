package repository

import (
	"errors"

	dbconection "library/domains/users/dbConection"
	"library/domains/users/entity"
)

func FindUserByEmail(email string) (user entity.User, err error) {
	db, err := dbconection.ConnectDB()
	if err != nil {
		return user, errors.New("Não foi possivel conectar ao banco de dados")
	}
	defer db.Close()

	rows, err := db.Query("SELECT BIN_TO_UUID(ID), Name, Email, Password, Owner FROM Users WHERE email = ?", email)
	if err != nil {
		return user, errors.New("Não foi possivel buscar usuário")
	}
	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Owner)
		if err != nil {
			return user, errors.New("Não foi possivel buscar usuário")
		}
	}

	return user, nil
}
