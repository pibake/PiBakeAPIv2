package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

var session *sql.DB

// JSONResponse struct: Response (tyipcally an error) back to the client in integer form
type JSONResponse struct {
	Response int `json:"response"`
}

// TempResponse struct: Response back for getting one temperature
type TempResponse struct {
	Temp Temp `json:"response"`
}

// TempsResponse struct: Response back for getting all of the temperatures
type TempsResponse struct {
	Temps []Temp `json:"response"`
}

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "PiBakeAPI v2, written and developed by Wyatt J. Miller, copyright 2019")
}

// GetTemps function: it's only a placeholder currently
func GetTemps(w http.ResponseWriter, r *http.Request) {
	var entry Temp
	var temps []Temp

	query, err := session.Query("SELECT * FROM `tempdata`")

	if err != nil {
		json.NewEncoder(w).Encode(JSONResponse{http.StatusBadRequest})
		return
	} 

	for query.Next() {
		err = query.Scan(&entry.UUID, &entry.Time, &entry.Date, &entry.TempF, &entry.TempC)
		temps = append(temps, entry)
	}

	if err != nil {
		json.NewEncoder(w).Encode(JSONResponse{http.StatusBadGateway})
	} else {
		json.NewEncoder(w).Encode(TempsResponse{temps})
	}
}

// GetTempOne function: WIP
func GetTempOne(w http.ResponseWriter, r *http.Request) {
	var entry Temp
	
	vars := mux.Vars(r)
	count := vars["id"]

	query, err := session.Query("SELECT * FROM `tempdata` WHERE `uuid` = ?", count)

	if err != nil {
		json.NewEncoder(w).Encode(JSONResponse{http.StatusBadRequest})
		return
	} 

	for query.Next() {
		err = query.Scan(&entry.UUID, &entry.Time, &entry.Date, &entry.TempF, &entry.TempC)
	}

	if err != nil {
		json.NewEncoder(w).Encode(JSONResponse{http.StatusBadGateway})
	} else {
		json.NewEncoder(w).Encode(TempResponse{entry})
	}
}

// DeleteTempOne function: it's only a placeholder currently
func DeleteTempOne(w http.ResponseWriter, r *http.Request) {
	// delete one code goes here
}
