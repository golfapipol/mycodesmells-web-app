package api

import (
	"db"
	"encoding/json"
	"net/http"
)

type API struct {
	DataSource db.DataSource
}

func (api API) GetCountries(w http.ResponseWriter, r *http.Request) {
	c, _ := api.DataSource.Countries()
	json.NewEncoder(w).Encode(&c)
}
