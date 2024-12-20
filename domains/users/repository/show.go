package repository

import (
	"errors"

	dbconection "library/domains/users/dbConection"
	"library/domains/users/entity"
)

func Show(ID string) (user entity.User, err error) {
	db, err := dbconection.ConnectDB()
	if err != nil {
		return user, errors.New("Não foi possivel conectar ao banco de dados")
	}

	query := "SELECT BIN_TO_UUID(ID), Name, Email, Owner, Phone FROM Users WHERE ID = UUID_TO_BIN(?)"

	stmt, err := db.Prepare(query)
	if err != nil {
		return user, errors.New("Não foi possivel buscar usuário")
	}
	defer stmt.Close()

	rows, err := stmt.Query(ID)
	if err != nil {
		return user, errors.New("Não foi possivel buscar usuário")
	}
	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.Owner,
			&user.Phone,
		)
		if err != nil {
			return user, errors.New("Não foi possivel buscar usuário")
		}

	}

	return user, nil
}
