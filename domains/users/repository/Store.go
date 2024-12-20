package repository

import (
	"errors"

	dbconection "library/domains/users/dbConection"
	"library/domains/users/entity"
	secutiry "library/utils/security"
)

func Store(user entity.User) error {
	db, err := dbconection.ConnectDB()
	if err != nil {
		return errors.New("Não foi possivel conectar ao banco de dados")
	}
	defer db.Close()

	hashPassword, err := secutiry.HashPassword(user.Password)
	if err != nil {
		return errors.New("Não foi possivel gerar hash da senha")
	}

	stmt, err := db.Prepare("INSERT INTO Users (ID, Name, Email, Password, Owner, Phone) VALUES (UUID_TO_BIN(?), ?, ?, ?, ?, ?)")
	if err != nil {
		return errors.New("Não foi possivel preparar a query")
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.ID, user.Name, user.Email, hashPassword, user.Owner, user.Phone)
	if err != nil {
		return errors.New("Não foi possivel executar a query")
	}

	return nil

}
