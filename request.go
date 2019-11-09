package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type DatabaseSession struct {
	session *sql.DB
}

type JSONResponse struct {
	Response string `json:"response"`
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

// CreateTempOne function that receieves a HTTP POST request
func CreateTempOne(w http.ResponseWriter, r *http.Request) {
	var entry Temp
	var dbQuery DatabaseSession
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&entry)

	if err != nil {
		log.Println("Couldn't decode the JSONs!")
		json.NewEncoder(w).Encode(JSONResponse{"400"})
	}

	query1, err := dbQuery.session.Query("SELECT * FROM `Pi` WHERE `uuid` = %v ", entry.UUID)

	if err != nil {
		log.Println("Database goofed: query1")
		json.NewEncoder(w).Encode(JSONResponse{"400"})
	}

	if err = query1.Scan(); err != nil {
		log.Println("New unit sending: inserting UUID for unit.")
		query2, err := dbQuery.session.Query("INSERT INTO `Pi` (uuid) VALUES %v", entry.UUID)

		defer query2.Close()

		if err != nil {
			log.Println("Database goofed: query2")
			json.NewEncoder(w).Encode(JSONResponse{"400"})
		}
	}

	query3, err := dbQuery.session.Query("INSERT INTO tempdata (uuid, date, temp_fahrenheit, temp_celsius) VALUES (%v, %v, %v, %v, %v)", entry.UUID,
		entry.Date,
		entry.Time,
		entry.TempF,
		entry.TempC)

	defer query3.Close()

	if err != nil {
		log.Println("Database goofed: query3")
		json.NewEncoder(w).Encode(JSONResponse{"400"})
	}
}

func DeleteTempOne(w http.ResponseWriter, r *http.Request) {
	// delete one code goes here
}
