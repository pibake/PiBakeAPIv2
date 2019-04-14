package main

type Temp struct {
	UUID  string `json:"uuid:omitempty"`
	Time  string `json:"time:omitempty"`
	Date  string `json:"date:omitempty"`
	TempF string `json:"temp_fahrenheit:omitempty"`
	TempC string `json:"temp_celsius:omitempty"`
}

var temp []Temp
