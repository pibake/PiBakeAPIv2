package main

import (
	"database/sql"
	"fmt"
	"net/http"
)

var session *sql.DB

type JSONResponse struct {
	Response int `json:"response"`
}

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "PiBakeAPI v2, written and developed by Wyatt J. Miller, copyright 2019")
}

func GetTemps(w http.ResponseWriter, r *http.Request) {
	// get all code goes here
}

func GetTempOne(w http.ResponseWriter, r *http.Request) {
	// get one code goes here
}

func DeleteTempOne(w http.ResponseWriter, r *http.Request) {
	// delete one code goes here
}
