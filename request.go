package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var session *sql.DB

type Temp struct {
	UUID  string `json:"uuid"`
	Time  string `json:"time"`
	Date  string `json:"date"`
	TempF int    `json:"temp_fahrenheit"`
	TempC int    `json:"temp_celsius"`
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
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&entry)

	if err != nil {
		log.Println("Couldn't decode the JSONs!")
		json.NewEncoder(w).Encode(JSONResponse{"400"})
	}

	query1, err := session.Exec("SELECT * FROM `pi` WHERE `uuid` = " + entry.UUID)

	log.Println(query1)

	if err != nil {
		log.Println("Database goofed: query1")
		json.NewEncoder(w).Encode(JSONResponse{"400"})
	}

	if err != nil {
		log.Println("New unit sending: inserting UUID for unit.")
		query2, err := session.Query("INSERT INTO `pi` (uuid) VALUES " + entry.UUID)

		defer query2.Close()

		if err != nil {
			log.Println("Database goofed: query2")
			json.NewEncoder(w).Encode(JSONResponse{"400"})
		}
	}

	query3, err := session.Query("INSERT INTO `tempdata` (uuid, date, temp_fahrenheit, temp_celsius) VALUES (%v, %v, %v, %v, %v)", entry.UUID,
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
