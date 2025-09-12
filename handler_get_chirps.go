package main

import (
	"net/http"

	"github.com/google/uuid"
)

func (cfg *apiConfig) getChirpsHandler(w http.ResponseWriter, r *http.Request) {

	dbChirps, err := cfg.dbQueries.GetChirps(r.Context())
	if err != nil {
		// insert error
		return
	}

	chirps := make([]Chirp, len(dbChirps))

	for i := range dbChirps {
		chirps[i].ID = dbChirps[i].ID
		chirps[i].CreatedAt = dbChirps[i].CreatedAt
		chirps[i].UpdatedAt = dbChirps[i].UpdatedAt
		chirps[i].Body = dbChirps[i].Body
		chirps[i].UserId = dbChirps[i].UserID
	}

	respondWithJSON(w, http.StatusOK, chirps)

}

// one chirp
func (cfg *apiConfig) getChirpHandler(w http.ResponseWriter, r *http.Request) {

	chirpIDStr := r.PathValue("chirpId")

	chirpID, err := uuid.Parse(chirpIDStr)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "invalid chirpId format", err)
	}

	dbChirp, err := cfg.dbQueries.GetChirp(r.Context(), chirpID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "chirp not found", err)
		return
	}

	chirp := Chirp{
		ID:        dbChirp.ID,
		CreatedAt: dbChirp.CreatedAt,
		UpdatedAt: dbChirp.UpdatedAt,
		Body:      dbChirp.Body,
		UserId:    dbChirp.UserID,
	}

	respondWithJSON(w, http.StatusOK, chirp)

}
