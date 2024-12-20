package pkgjwt

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

func Auth(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	authHead := strings.Split(r.Header.Get("Authorization"), "Bearer")
	if len(authHead) != 2 {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Malformed Token"))

	} else {
		jwtToken := authHead[1]
		token, _ := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, errors.New("Falha na validação")
			}

			return []byte("JWT_SECRET"), nil
		})

		if token != nil && token.Claims == nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Não autorizado"))
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			ctx := context.WithValue(r.Context(), "props", claims)
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Não autorizado"))
			return
		}
	}
}
