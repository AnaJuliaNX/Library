package users

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"

	"library/domains/users/entity"
	"library/domains/users/presenter"
	respository "library/domains/users/repository"
	response "library/utils/responseError"
	utils "library/utils/security"
)

func Login(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		response.ErrorMessage(w, http.StatusUnprocessableEntity, errors.New("Erro ao ler os dados"))
		return
	}

	var userLogin entity.User
	err = json.Unmarshal(data, &userLogin)
	if err != nil {
		response.ErrorMessage(w, http.StatusUnprocessableEntity, errors.New("Erro ao converter os dados"))
		return
	}

	user, err := respository.FindUserByEmail(userLogin.Email)
	if err != nil {
		response.ErrorMessage(w, http.StatusUnprocessableEntity, errors.New("Usuário não encontrado"))
		return
	}

	log.Println(user.Password, userLogin.Password)

	err = utils.NewPass().ComaprePassword(user.Password, userLogin.Password)
	if err != nil {
		log.Println(err)
		response.ErrorMessage(w, http.StatusUnprocessableEntity, errors.New("Email ou senha não confere"))
		return
	}

	token, expired, err := user.GenerateToken()
	if err != nil {
		response.ErrorMessage(w, http.StatusUnprocessableEntity, errors.New("Erro ao gerar token"))
		return
	}

	var respUser presenter.User
	respUser.ID = user.ID
	respUser.Name = user.Name
	respUser.Email = user.Email
	respUser.Phone = user.Phone
	respUser.Owner = user.Owner
	respUser.SetMeta(token, "Bearer", expired)

	response.Success(w, http.StatusOK, respUser)
}
