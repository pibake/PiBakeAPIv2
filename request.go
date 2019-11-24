package main

import (
	"database/sql"
	"fmt"
	"net/http"
)

var session *sql.DB

// JSONResponse struct: Response back to the client in integer form
type JSONResponse struct {
	Response int `json:"response"`
}

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "PiBakeAPI v2, written and developed by Wyatt J. Miller, copyright 2019")
}

// GetTemps function: it's only a placeholder currently
func GetTemps(w http.ResponseWriter, r *http.Request) {
	// get all code goes here
}

// GetTempOne function: it's only a placeholder currently
func GetTempOne(w http.ResponseWriter, r *http.Request) {
	// get one code goes here
}

// DeleteTempOne function: it's only a placeholder currently
func DeleteTempOne(w http.ResponseWriter, r *http.Request) {
	// delete one code goes here
}
