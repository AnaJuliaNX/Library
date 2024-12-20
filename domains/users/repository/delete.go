package repository

import (
	"errors"
	dbconection "library/domains/users/dbConection"
)

func DeleteUser(userId string) error {
	db, err := dbconection.ConnectDB()
	if err != nil {
		return errors.New("Não foi possivel conectar ao banco de dados")
	}
	defer db.Close()

	stm, err := db.Prepare("DELETE FROM Users WHERE ID = UUID_TO_BIN(?)")
	if err != nil {
		return errors.New("Erro ao deletar usuário")
	}

	_, err = stm.Exec(userId)
	if err != nil {
		return errors.New("Erro ao deletar usuário")
	}

	return nil

}
