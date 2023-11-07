package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	dat, err := json.Marshal(payload)
	if err != nil {
		fmt.Printf("Failed to marshal payload %v", payload)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(dat)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	if code > 499 {
		fmt.Printf("Responding with 5XX error: %s", message)
		return
	}

	type errResponse struct {
		Error string `json:"error"`
	}

	respondWithJson(w, code, errResponse{Error: message})
}
