package test

import (
	"db"
)

type FakeDataSource struct{}

func (FakeDataSource) Countries() ([]db.Country, error) {
	c := []db.Country{
		db.Country{Code: "PL", Name: "Poland", Capital: "Warsaw"},
		db.Country{Code: "USA", Name: "USA", Capital: "Washington"},
	}
	return c, nil
}
