package utils

import (
	"encoding/json"
	"errors"
	"io"
	"library/domains/users/entity"
	"net/http"
)

func ExtractBody(r *http.Request, data *entity.User) error {
	bodyData, err := io.ReadAll(r.Body)
	if err != nil {
		return errors.New("Não foi possivel extrair o body da requisição")
	}

	err = json.Unmarshal(bodyData, &data)
	if err != nil {
		return errors.New("Não foi possivel converter o body da requisição")
	}

	return nil
}
