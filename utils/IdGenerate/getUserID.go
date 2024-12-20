package idgenerate

import (
	"errors"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

func GetUserID(r *http.Request) (string, error) {
	props := r.Context().Value("props")
	data := props.(jwt.MapClaims)

	log.Println(data["userid"])
	if data["userid"] == nil || (data["userid"] != nil && data["userid"].(string) == "") {
		return "", errors.New("Usuário não encontrado")
	}

	return data["userid"].(string), nil
}
