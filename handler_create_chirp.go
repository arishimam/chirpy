package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/arishimam/chirpy/internal/database"
	"github.com/google/uuid"
)

type Chirp struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Body      string    `json:"body"`
	UserId    uuid.UUID `json:"user_id"`
}

func (cfg *apiConfig) createChirpHandler(w http.ResponseWriter, r *http.Request) {

	type parameters struct {
		Body   string    `json:"body"`
		UserId uuid.UUID `json:"user_id"`
	}

	params := parameters{}

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		//http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	body, err := validateChirp(params.Body)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error(), err)
		return
	}

	chirp, err := cfg.dbQueries.CreateChirp(r.Context(), database.CreateChirpParams{
		Body:   body,
		UserID: params.UserId,
	})

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Couldn't create chirp in db", nil)
		return

	}

	respondWithJSON(w, http.StatusCreated, Chirp{
		ID:        chirp.ID,
		CreatedAt: chirp.CreatedAt,
		UpdatedAt: chirp.CreatedAt,
		Body:      chirp.Body,
		UserId:    chirp.UserID,
	})

}

func validateChirp(body string) (string, error) {
	const maxChirpLength = 140

	if len(body) > maxChirpLength {
		return "", errors.New("Chirp is too long")
	}

	cleanedBody := checkProfane(body)

	return cleanedBody, nil

}

func checkProfane(msg string) string {
	bannedMap := map[string]bool{
		"kerfuffle": true,
		"sharbert":  true,
		"fornax":    true,
	}

	words := strings.Fields(msg)

	for i, w := range words {
		if bannedMap[strings.ToLower(w)] {
			words[i] = "****"
		}
	}

	cleanedMsg := strings.Join(words, " ")

	return cleanedMsg
}
