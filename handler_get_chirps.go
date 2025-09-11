package main

import "net/http"

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
/*
func (cfg *apiConfig) getChirpHandler(w http.ResponseWriter, r *http.Request) {

	dbChirp, err := cfg.dbQueries.GetChirp(r.Context())
	if err != nil {
		// insert error
		return
	}

	respondWithJSON(w, http.StatusOK, dbChirp)

}
*/
