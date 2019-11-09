package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	var session DatabaseSession
	db := NewSQLDB("wyatt", "wyattisawesome", "thepibake.com", "3306", "PiBake_Test")
	session.session = db.OpenSQL()
	router := mux.NewRouter()

	router.HandleFunc("/api", root).Methods("GET")
	router.HandleFunc("/api/temps", GetTemps).Methods("GET")
	router.HandleFunc("/api/temps/{id}", GetTempOne).Methods("GET")
	router.HandleFunc("/api/temps/create", CreateTempOne).Methods("POST")
	router.HandleFunc("/api/temps/{id}", DeleteTempOne).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))

}
