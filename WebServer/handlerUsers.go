package main

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/shri771/Go/WebServer/internal/auth"
	"github.com/shri771/Go/WebServer/internal/database"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Email     string    `json:"email"`
}

func (cfg *apiConfig) handlerUsers(w http.ResponseWriter, r *http.Request) {
	type users struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	type response struct {
		User
	}
	//Decode Json
	user := users{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Erro with Decoding:", err)
		return
	}
	defer r.Body.Close()

	hashedPassword, err := auth.HashPassword(user.Password)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error with hashing pswd", err)
		return
	}

	addedUser, err := cfg.db.CreateUser(context.Background(), database.CreateUserParams{
		CreatedAt:      time.Now().UTC(),
		UpdatedAt:      time.Now().UTC(),
		Email:          user.Email,
		HashedPassword: hashedPassword,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error with Adding user:", err)
		return
	}

	//Encode Json
	respondWithJSON(w, http.StatusCreated, response{
		User: User{
			ID:        addedUser.ID,
			CreatedAt: addedUser.CreatedAt,
			UpdatedAt: addedUser.UpdatedAt,
			Email:     addedUser.Email,
		},
	})
}

func (cfg *apiConfig) handlerLogin(w http.ResponseWriter, r *http.Request) {
	type users struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// Decode Data
	loginCredintial := users{}

	err := json.NewDecoder(r.Body).Decode(&loginCredintial)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Error with decoding json", err)
		return
	}
	defer r.Body.Close()

	// Query database
	login, err := cfg.db.GetUserByEmail(context.Background(), loginCredintial.Email)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Error while quering database", err)
		return
	}

	// Verify user
	err = auth.CheckPasswordHash(loginCredintial.Password, login.HashedPassword)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Error while quering database", err)
		return
	}

	respondWithJSON(w, http.StatusOK, User{
		ID:        login.ID,
		CreatedAt: login.CreatedAt,
		UpdatedAt: login.UpdatedAt,
		Email:     login.Email,
	})

}
