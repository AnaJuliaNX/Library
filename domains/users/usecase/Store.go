package users

import (
	"errors"
	"net/http"

	entity "library/domains/users/entity"
	"library/domains/users/presenter"
	"library/domains/users/repository"
	user "library/domains/users/repository"
	utils "library/utils"
	utilsID "library/utils/IdGenerate"
	response "library/utils/responseError"
)

func Store(w http.ResponseWriter, r *http.Request) {
	var User entity.User
	err := utils.ExtractBody(r, &User)
	if err != nil {
		response.ErrorMessage(w, http.StatusUnprocessableEntity, errors.New("Não foi possivel extrair dados da requisição"))
		return
	}

	userEmail, err := user.FindUserByEmail(User.Email)
	if err != nil {
		response.ErrorMessage(w, http.StatusUnprocessableEntity, errors.New("Não foi possivel buscar usuário"))
		return
	}

	if userEmail.Email != "" {
		response.ErrorMessage(w, http.StatusUnprocessableEntity, errors.New("Usuário já cadastrado"))
		return
	}

	userID := utilsID.NewID().String()
	User.ID = userID

	err = repository.Store(User)
	if err != nil {
		response.ErrorMessage(w, http.StatusUnprocessableEntity, errors.New("Não foi possivel cadastrar usuário"))
		return
	}

	token, exp, err := User.GenerateToken()
	if err != nil {
		response.ErrorMessage(w, http.StatusUnprocessableEntity, errors.New("Não foi possivel gerar token"))
		return
	}

	var presenterUser presenter.User
	presenterUser.ID = User.ID
	presenterUser.Name = User.Name
	presenterUser.Email = User.Email
	presenterUser.Phone = User.Phone
	presenterUser.Owner = User.Owner

	presenterUser.SetMeta(token, "Bearer", exp)

	utils.Sucess(w, http.StatusCreated, presenterUser)
}
