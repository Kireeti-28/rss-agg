package main

import "net/http"


func readinessHanldler(w http.ResponseWriter, r *http.Request) {
	respondWithJson(w, 200, struct{}{})
}
