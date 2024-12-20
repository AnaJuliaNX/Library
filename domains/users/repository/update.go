package repository

import (
	"errors"
	"fmt"
	dbconection "library/domains/users/dbConection"
	"library/domains/users/entity"
	"strings"
)

func Update(user entity.User, userID string) error {
	db, err := dbconection.ConnectDB()
	if err != nil {
		return errors.New("Não foi possivel conectar ao banco de dados")
	}
	defer db.Close()

	fields := []string{}
	values := []interface{}{}

	if user.Name != "" {
		fields = append(fields, "Name = ?")
		values = append(values, user.Name)
	}

	if user.Email != "" {
		fields = append(fields, "Email = ?")
		values = append(values, user.Email)
	}

	if user.Phone != "" {
		fields = append(fields, "Phone = ?")
		values = append(values, user.Phone)
	}

	if len(fields) == 0 {
		return errors.New("Não foi possivel atualizar usuário")
	}

	query := fmt.Sprintf("UPDATE Users SET %s WHERE ID = UUID_TO_BIN(?)", strings.Join(fields, ", "))
	values = append(values, userID)

	stmt, err := db.Prepare(query)
	if err != nil {
		return errors.New("Não foi possivel atualizar usuário")
	}
	defer stmt.Close()

	_, err = stmt.Exec(values...)
	if err != nil {
		return errors.New("Não foi possivel atualizar usuário")
	}

	return nil
}
