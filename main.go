package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// placeholder creds, don't expect them to get you anywhere
	db := NewSQLDB("wymillerlinux", "123456789", "localhost", "3306", "PiBakeAPIv2")
	session = db.OpenSQL()
	defer session.Close()

	router := mux.NewRouter()

	router.HandleFunc("/api", root).Methods("GET")
	router.HandleFunc("/api/temps", GetTemps).Methods("GET") // placeholder
	router.HandleFunc("/api/temps/{id}", GetTempOne).Methods("GET") // placeholder
	router.HandleFunc("/api/temps/create", CreateTempOne).Methods("POST")
	router.HandleFunc("/api/temps/{id}", DeleteTempOne).Methods("DELETE") // placeholder

	fmt.Println("Starting PiBakeAPI version 2...")

	log.Fatal(http.ListenAndServe(":8000", router))

}
