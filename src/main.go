package main

import (
	"encoding/json"
	"net/http"
)

func GetCountries(w http.ResponseWriter, r *http.Request) {
	c := []Country{
		Country{Code: "PL", Name: "Poland", Capital: "Warsaw"},
		Country{Code: "USA", Name: "USA", Capital: "Washington"},
	}
	json.NewEncoder(w).Encode(&c)
}
