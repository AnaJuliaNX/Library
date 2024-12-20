package users

import (
	"errors"
	"net/http"

	"github.com/gorilla/mux"

	"library/domains/users/presenter"
	"library/domains/users/repository"
	response "library/utils/responseError"
)

func Show(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	if params["id"] == "" {
		response.ErrorMessage(w, http.StatusUnprocessableEntity, errors.New("Usuário não encontrado"))
		return
	}

	user, err := repository.Show(params["id"])
	if err != nil {
		response.ErrorMessage(w, http.StatusUnprocessableEntity, err)
		return
	}

	var respUser presenter.Users
	if user.ID != "" {
		respUser.ID = user.ID
		respUser.Name = user.Name
		respUser.Email = user.Email
		respUser.Phone = user.Phone
		respUser.Owner = user.Owner
	}

	response.Success(w, http.StatusOK, respUser)
}
