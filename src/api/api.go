package api

import (
	"encoding/json"
	"net/http"
)

type API struct {
	DataSource db.DataSource
}

func (api API) GetCountries(w http.ResponseWriter, r *http.Request) {
	c := []Country{
		Country{Code: "PL", Name: "Poland", Capital: "Warsaw"},
		Country{Code: "USA", Name: "USA", Capital: "Washington"},
	}
	json.NewEncoder(w).Encode(&c)
}
