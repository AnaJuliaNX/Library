package utils

import (
	"net/http"
	"strconv"
	"unicode"
)

func ExtrairNumeros(texto string) string {
	var numeros string

	for _, char := range texto {
		if unicode.IsDigit(char) {
			numeros += string(char)
		}
	}

	return numeros
}

func GetPageInQuery(r *http.Request) (string, int) {
	var pageStr string = ExtrairNumeros(r.URL.Query().Get("page"))
	if pageStr != "" {
		page, err := strconv.Atoi(pageStr)
		if err != nil {
			return "1", 1
		}

		return pageStr, page
	}

	return "1", 1
}
