package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Chirp struct {
	Body  string `json: "body"`
	Valid bool   `json: "valid"`
}

func (cfg *apiConfig) validateHandler(w http.ResponseWriter, r *http.Request) {
	var chirp Chirp

	err := json.NewDecoder(r.Body).Decode(&chirp)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
	}
	defer r.Body.Close()

	if len(chirp.Body) > 140 {
		http.Error(w, "Chirp is too long", http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf(`{"valid": true}`)))
	}

	fmt.Println("Chirp Body: ", chirp.Body)

	fmt.Println("Chirp Valid : ", chirp.Valid)

}
