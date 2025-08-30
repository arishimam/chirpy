package main

import (
	"fmt"
	"net/http"
)

func (cfg *apiConfig) resetHandler(w http.ResponseWriter, r *http.Request) {

	if cfg.platform != "dev" {
		// respond with 403 forbidden
		respondWithError(w, http.StatusForbidden, "PLATFORM environment variable is not set to dev.", nil)
		return
	}

	cfg.dbQueries.DeleteUsers(r.Context())

	cfg.fileserverHits.Store(0)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("All users have been deleted and hits have been reset to 0!")))
}
