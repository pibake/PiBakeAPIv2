package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// CreateTempOne function that receieves a HTTP POST request and stores data in database.
// Analogous to the import script we had earlier in the project
func CreateTempOne(w http.ResponseWriter, r *http.Request) {
	var entry Temp
	var count int
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&entry)

	if err != nil {
		log.Println("Couldn't decode the JSONs!")
		json.NewEncoder(w).Encode(JSONResponse{"400"})
	}

	query1, err := session.Query("SELECT COUNT(*) FROM `pi` WHERE `uuid` = ?", entry.UUID)

	defer query1.Close()

	if err != nil {
		log.Println("Database goofed: query1")
		json.NewEncoder(w).Encode(JSONResponse{"400"})
	}

	// check to see if any rows match the supplied UUID
	for query1.Next() {
		if err := query1.Scan(&count) ; err != nil {
			log.Println("Scanning the database did a thing scanning the UUIDs")
			json.NewEncoder(w).Encode(JSONResponse{"500"})
		}
	}

	// if the count from the scan is 0, insert the UUID in its place
	if count == 0 {
		query2, err := session.Query("INSERT INTO `pi` (uuid) VALUES (?)", entry.UUID)

		defer query2.Close()

		if err != nil {
			log.Println("Database goofed: query2")
			json.NewEncoder(w).Encode(JSONResponse{"400"})
		}
	}

	query3, err := session.Query("INSERT INTO `tempdata` (uuid, date, time, temp_fahrenheit, temp_celsius) VALUES (?, ?, ?, ?, ?)", entry.UUID, entry.Date, entry.Time, entry.TempF, entry.TempC)

	defer query3.Close()

	if err != nil {
		log.Println("Database goofed: query3")
		json.NewEncoder(w).Encode(JSONResponse{"400"})
	}
}