package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/arishimam/chirpy/internal/database"
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
		Password string `json:"password"`
		Email    string `json:"email"`
	}

	parameters := Parameters{}

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&parameters)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	hashedPass := parameters.Password
	sqlEmail := toNullString(parameters.Email)

	createUserParams := database.CreateUserParams{
		Email:          sqlEmail,
		HashedPassword: hashedPass,
	}

	user, err := cfg.dbQueries.CreateUser(r.Context(), createUserParams)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	userMapped := User{
		user.ID,
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
