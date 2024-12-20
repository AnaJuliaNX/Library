package users

import (
	"errors"
	"net/http"

	"github.com/gorilla/mux"

	repository "library/domains/users/repository"
	response "library/utils/responseError"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	if params["id"] == "" {
		response.ErrorMessage(w, http.StatusBadRequest, errors.New("Usuário não encontrado"))
	}

	err := repository.DeleteUser(params["id"])
	if err != nil {
		response.ErrorMessage(w, http.StatusBadRequest, errors.New("Erro ao deletar usuário"))
		return
	}

	response.Success(w, http.StatusOK, "Usuário deletado com sucesso")
}
