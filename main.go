package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	db := NewSQLDB("wyatt", "wyattisawesome", "localhost", "3306", "PiData_Test")
	session = db.OpenSQL()
	defer session.Close()

	router := mux.NewRouter()

	router.HandleFunc("/api", root).Methods("GET")
	router.HandleFunc("/api/temps", GetTemps).Methods("GET")
	router.HandleFunc("/api/temps/{id}", GetTempOne).Methods("GET")
	router.HandleFunc("/api/temps/create", CreateTempOne).Methods("POST")
	router.HandleFunc("/api/temps/{id}", DeleteTempOne).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))

}
