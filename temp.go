package main

// Temp struct, all the JSON fields, all neat and tidy
type Temp struct {
	UUID  string `json:"uuid"`
	Time  string `json:"time"`
	Date  string `json:"date"`
	TempF int    `json:"temp_fahrenheit"`
	TempC int    `json:"temp_celsius"`
}