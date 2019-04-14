package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// entry point ??
func main() {
	router := mux.NewRouter()

	router.HandleFunc("/temps", GetTemps).Methods("GET")
	router.HandleFunc("/temps/{id}", GetTempOne).Methods("GET")
	router.HandleFunc("/temps/{id}", CreateTempOne).Methods("POST")
	router.HandleFunc("/temps/{id}", DeleteTempOne).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))

}
