package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Email     string    `json:"email"`
}

func (cfg *apiConfig) createUserHandler(w http.ResponseWriter, r *http.Request) {

	type Parameters struct {
		Email string `json:"email"`
	}

	parameters := Parameters{}

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&parameters)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	sqlEmail := toNullString(parameters.Email)

	user, err := cfg.dbQueries.CreateUser(r.Context(), sqlEmail)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	userMapped := User{
		user.ID.UUID,
		user.CreatedAt.Time,
		user.UpdatedAt.Time,
		user.Email.String,
	}

	respondWithJSON(w, http.StatusCreated, userMapped)

}

func toNullString(s string) sql.NullString {
	if s == "" {
		return sql.NullString{String: "", Valid: false}
	}

	return sql.NullString{String: s, Valid: true}

}

func fromNullString(s sql.NullString) string {
	if s.Valid {
		return s.String
	}

	return "ERROR"

}
