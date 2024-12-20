package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	user "library/domains/users/usecase"
)

func main() {

	router := mux.NewRouter()
	router.Use(RequestExecuted)

	//validRouter := router.PathPrefix("/v1").Subrouter()
	//validRouter.Use(jwt.Auth)

	router.HandleFunc("/v1/users/login", user.Login).Methods(http.MethodPost)
	router.HandleFunc("/v1/users/store", user.Store).Methods(http.MethodPost)
	router.HandleFunc("/v1/users/index", user.Index).Methods(http.MethodGet)
	router.HandleFunc("/v1/users/show/{id}", user.Show).Methods(http.MethodGet)
	router.HandleFunc("/v1/users/update", user.Update).Methods(http.MethodPut)
	router.HandleFunc("/v1/users/delete/{id}", user.Delete).Methods(http.MethodDelete)

	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func RequestExecuted(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}
