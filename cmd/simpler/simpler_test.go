package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/test", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.Handler(simpleHandler{})
	handler.ServeHTTP(rr, req)

	//Check response code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	//Check body
	decoder := json.NewDecoder(rr.Body)
	response := struct {
		Path      string `json:"path"`
		Method    string `json:"method"`
		Pid       int    `json:"pid"`
		Hostname  string `json:"hostname"`
		Timestamp int64  `json:"timestamp"`
	}{}
	err = decoder.Decode(&response)
	if err != nil {
		t.Errorf("handler returned unexpected json error: %v", err)
	}

	if response.Path != "/test" {
		t.Errorf("handle returned unexpected path: got %v want /test", response.Path)
	}

	if response.Method != "GET" {
		t.Errorf("handle returned unexpected method : got %v want GET", response.Method)
	}

	if rr.Header().Get("Content-Type") != "application/json" {
		t.Errorf("handle returned unexpected content type: got %v want application/json", rr.Header().Get("Content-Type"))
	}
}
