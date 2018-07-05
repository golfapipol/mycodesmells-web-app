package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func init() {
	http.HandleFunc("/countries", GetCountries)
}

func TestGetCountries(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/countries", nil)
	rw := httptest.NewRecorder()
	http.DefaultServeMux.ServeHTTP(rw, req)

	if rw.Code != 200 {
		t.Fatalf("expected 200 response code, but got: %v\n", rw.Code)
	}
	var c []Country
	err := json.NewDecoder(rw.Body).Decode(&c)
	if err != nil {
		t.Fatal("Failed to decode countries list")
	}

	if len(c) != 2 {
		t.Fatalf("Expected length of countries list to equal 2, not: %v\n", len(c))
	}
	t.Log("Expected length of countries list to equal 2", len(c))

	pl := Country{Code: "PL", Name: "Poland", Capital: "Warsaw"}
	if c[0] != pl {
		t.Fatalf("Expected first country to be Poland, not: %v\n", c[0])
	}
	t.Log("Expected first country to be Poland")

	usa := Country{Code: "USA", Name: "USA", Capital: "Washington"}
	if c[1] != usa {
		t.Fatalf("Expected first country to be USA, not: %v\n", c[1])
	}
}
