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
	ID           uuid.UUID `json:"id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Email        string    `json:"email"`
	Token        string    `json:"token"`
	RefreshToken string    `json:"refresh_token"`
	Member       bool      `json:"is_chirpy_red"`
}

func (cfg *apiConfig) handlerUsersCreate(w http.ResponseWriter, r *http.Request) {
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

	// Hash Password and Add User to Database
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

	// Return Json
	respondWithJSON(w, http.StatusCreated, response{
		User: User{
			ID:        addedUser.ID,
			CreatedAt: addedUser.CreatedAt,
			UpdatedAt: addedUser.UpdatedAt,
			Email:     addedUser.Email,
			Member:    addedUser.IsChipryRed,
		},
	})
}
