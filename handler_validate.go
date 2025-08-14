package main

import (
	"encoding/json"
	"net/http"
)

func (cfg *apiConfig) validateChirpHandler(w http.ResponseWriter, r *http.Request) {

	type parameters struct {
		Body string `json:"body"`
	}

	type returnVals struct {
		Valid bool `json:"valid"`
	}

	params := parameters{}

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		//http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	const maxChirpLength = 140

	if len(params.Body) > maxChirpLength {
		respondWithError(w, http.StatusBadRequest, "Chirp is too long", nil)
		//http.Error(w, "Chirp is too long", http.StatusBadRequest)
		return
	}

	respondWithJSON(w, http.StatusOK, returnVals{
		Valid: true,
	})

}
