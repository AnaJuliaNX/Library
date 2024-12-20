package pkgjwt

import (
	"errors"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Criando o token
func GenerateToken(
	userID string,
	email string,
	isOwner bool,
	DateCreated string,
) (tokenSigned string, expiration int64, err error) {
	expiration = time.Now().Add(time.Hour * 24 * 30).Unix()

	permissions := jwt.MapClaims{}
	permissions["exp"] = expiration
	permissions["userid"] = userID
	permissions["email"] = email
	permissions["isOwner"] = isOwner
	permissions["DateCreated"] = DateCreated

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)
	tokenSigned, err = token.SignedString([]byte("secret"))
	if err != nil {
		return "", 0, err
	}

	return tokenSigned, expiration, nil
}

// Validando se é valido
func ValidateToken(r *http.Request) error {
	tokenString := tokenExtract(r)
	token, err := jwt.Parse(tokenString, verifiToken)
	if err != nil {
		return err
	}

	_, ok := token.Claims.(jwt.MapClaims)
	if !ok && token.Valid {
		return nil
	}

	return errors.New("Token inválido")
}

func verifiToken(token *jwt.Token) (interface{}, error) {
	_, ok := token.Method.(*jwt.SigningMethodHMAC)
	if !ok {
		return nil, errors.New("Token inválido")
	}

	return os.Getenv("JWT_SECRET"), nil
}

func tokenExtract(r *http.Request) string {
	token := r.Header.Get("Authorization")
	split := strings.Split(token, " ")
	if len(split) == 2 {
		return split[1]
	}
	return ""
}

// Pegando o token

func ExtractToken(r *http.Request) (claims *jwt.MapClaims, err error) {
	tokenStr := tokenExtract(r)
	if tokenStr == "" {
		return nil, errors.New("Token não encontrado")
	}

	token, err := jwt.Parse(tokenStr, verifiToken)
	if err != nil {
		return nil, err
	}

	permissions, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return &permissions, nil
	}

	return nil, errors.New("Token inválido")

}
