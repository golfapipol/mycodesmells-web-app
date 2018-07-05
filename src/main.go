package main

import (
	_api "api"
	"db"
	"net/http"
)

type Country struct {
	Code, Name, Capital string
}

func main() {
	datasource, err := db.NewMongo("localhost", "go_countries")
	defer datasource.Session.Close()
	if err != nil {
		panic(err.Error())
	}
	api := _api.API{
		DataSource: datasource,
	}
	http.HandleFunc("/countries", api.GetCountries)
	http.ListenAndServe(":3000", nil)
}
