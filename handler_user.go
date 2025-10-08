package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/arishimam/chirpy/internal/auth"
	"github.com/arishimam/chirpy/internal/database"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Email     string    `json:"email"`
}

type Parameters struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (cfg *apiConfig) createUserHandler(w http.ResponseWriter, r *http.Request) {

	parameters := Parameters{}

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&parameters)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	hashedPass, err := auth.HashPassword(parameters.Password)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't hash the password", err)
		return
	}

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

func (cfg *apiConfig) loginUserHandler(w http.ResponseWriter, r *http.Request) {

	parameters := Parameters{}

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&parameters)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	email := toNullString(parameters.Email)
	user, err := cfg.dbQueries.LookupUserFromEmail(r.Context(), email)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "unable to lookup user from email", err)
		return
	}

	match, err := auth.CheckPasswordHash(parameters.Password, user.HashedPassword)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "error calling CheckPsswordHash", err)
		return
	}

	if match == false {
		respondWithJSON(w, http.StatusUnauthorized, "incorrect email or password.")
		return
	}

	userMapped := User{
		user.ID,
		user.CreatedAt.Time,
		user.UpdatedAt.Time,
		user.Email.String,
	}
	respondWithJSON(w, http.StatusOK, userMapped)

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
