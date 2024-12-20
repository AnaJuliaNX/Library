package responseerrorgo

import (
	"encoding/json"
	"log"
	"net/http"
)

func ErrorMessage(w http.ResponseWriter, statusCode int, err error) {
	JSONform(w, statusCode, prepare(struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Errors  error  `json:"errors"`
	}{
		Code:    statusCode,
		Message: err.Error(),
	}))
}

func JSONform(w http.ResponseWriter, statusCode int, data interface{}) {
	var err error
	w.WriteHeader(statusCode)
	if data != nil {
		err = json.NewEncoder(w).Encode(data)
		if err != nil {
			log.Fatal(err)
		}
	}
}

type dataString struct {
	Data interface{} `json:"data"`
}

func prepare(data interface{}) dataString {
	return dataString{Data: data}
}

// JSON prepara e responde com um json
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	var err error
	w.WriteHeader(statusCode)
	if data != nil {
		err = json.NewEncoder(w).Encode(data)
		if err != nil {
			log.Fatal(err)
		}
	}
}

// Prepara mensagem padrão de sucesso
func Success(w http.ResponseWriter, statusCode int, data interface{}) {
	JSON(w, statusCode, prepare(data))
}
