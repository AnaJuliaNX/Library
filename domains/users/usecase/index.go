package users

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"library/domains/users/presenter"
	repository "library/domains/users/repository"
	response "library/utils/responseError"
)

func Index(w http.ResponseWriter, r *http.Request) {
	search := r.URL.Query().Get("search")
	page := r.URL.Query().Get("page")
	atualPage, err := strconv.Atoi(page)
	if err != nil {
		atualPage = 1
	}

	limit := 15
	offset := limit * (atualPage - 1)

	users, total, err := repository.Index(search, int64(offset), int64(limit))
	if err != nil {
		response.ErrorMessage(w, http.StatusUnprocessableEntity, errors.New("Não foi possível buscar usuários"))
		return
	}

	var userPresenters []presenter.Users
	for _, user := range users {
		userPresenters = append(userPresenters, presenter.Users{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
			Phone: user.Phone,
			Owner: user.Owner,
		})
	}

	resp := map[string]interface{}{
		"data":       userPresenters,
		"total":      total,
		"page":       page,
		"totalPages": (total + int(limit) - 1) / int(limit),
	}

	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(resp); err != nil {
		response.ErrorMessage(w, http.StatusUnprocessableEntity, errors.New("Não foi possivel retornar os dados"))
		return
	}
}
