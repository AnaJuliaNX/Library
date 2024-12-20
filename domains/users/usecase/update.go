package users

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"library/domains/users/entity"
	repository "library/domains/users/repository"
	utils "library/utils/IdGenerate"
	response "library/utils/responseError"
)

func Update(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		response.ErrorMessage(w, http.StatusUnprocessableEntity, errors.New("Não foi possivel ler os dados"))
		return
	}

	var user entity.User
	err = json.Unmarshal(data, &user)
	if err != nil {
		response.ErrorMessage(w, http.StatusUnprocessableEntity, errors.New("Não foi possivel decodificar os dados"))
		return
	}

	userID, err := utils.GetUserID(r)
	if err != nil {
		response.ErrorMessage(w, http.StatusUnprocessableEntity, err)
		return
	}

	err = repository.Update(user, userID)
	if err != nil {
		response.ErrorMessage(w, http.StatusUnprocessableEntity, err)
		return
	}

	response.Success(w, http.StatusOK, "Usuário atualizado com sucesso")
}
