package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/shri771/Go/WebServer/internal/auth"
	"github.com/shri771/Go/WebServer/internal/database"
)

func (cfg *apiConfig) handlerUsersUpdate(w http.ResponseWriter, r *http.Request) {
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
		respondWithError(w, http.StatusUnauthorized, "Error with hashing pswd", err)
		return
	}

	userID, err := cfg.getUserIDJWT(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Could not retrive token from auth header", err)
		return
	}

	updatedUser, err := cfg.db.UpdateUser(context.Background(), database.UpdateUserParams{
		Email:          user.Email,
		HashedPassword: hashedPassword,
		ID:             userID,
	})

	respondWithJSON(w, http.StatusOK, response{
		User: User{
			ID:        updatedUser.ID,
			CreatedAt: updatedUser.CreatedAt,
			UpdatedAt: updatedUser.UpdatedAt,
			Email:     updatedUser.Email,
			Member:    updatedUser.IsChipryRed,
		},
	})

}
