package utils

import (
	"encoding/json"
	"log"
	"math"
	"net/http"

	entity "library/utils/entity"
)

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

func Sucess(w http.ResponseWriter, status int, data interface{}) {
	JSON(w, status, prepare(data))
}

func Message(w http.ResponseWriter, statusCode int, data interface{}) {
	JSON(w, statusCode, prepare(struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}{
		Code:    statusCode,
		Message: data.(string),
	}))
}

type dataSTR struct {
	Data interface{} `json:"data"`
}

// prepare encapsula resultados dentro do data tag
func prepare(data interface{}) interface{} {
	return dataSTR{Data: data}
}

func Pagination(total int64, returnData interface{}) entity.Response {
	var meta entity.Meta
	meta.CurrentPage = 1
	meta.Total = total

	totalPages := total / 15
	if totalPages > 0 {
		meta.TotalPages = int64(math.Round(float64(totalPages)))
	}
	if meta.Total != 0 && meta.TotalPages == 0 {
		meta.TotalPages = 1
	}

	var response entity.Response
	response.Data = returnData
	response.Meta = meta

	return response
}
