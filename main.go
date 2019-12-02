package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// placeholder creds, don't expect them to get you anywhere
	db := NewSQLDB("pibake", "awesomestuff", "localhost", "3306", "PiBakeApi")
	session = db.OpenSQL()
	defer session.Close()

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/api", root).Methods("GET")
	router.HandleFunc("/api/temp", GetTemps).Methods("GET")
	router.HandleFunc("/api/temp/{id}", GetTempOne).Methods("GET")
	router.HandleFunc("/api/temp/create", CreateTempOne).Methods("POST")
	router.HandleFunc("/api/temp/{id}", DeleteTempOne).Methods("DELETE") // placeholder

	fmt.Println("Starting PiBakeAPI version 2...")

	log.Fatal(http.ListenAndServe(":8000", router))

}
