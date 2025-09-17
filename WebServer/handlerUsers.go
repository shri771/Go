package main

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/shri771/Go/WebServer/internal/database"
)

func (cfg *apiConfig) handlerUsers(w http.ResponseWriter, r *http.Request) {
	type users struct {
		Email string `json:"email"`
	}

	//Decode Json
	user := users{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Erro with Decoding:", err)
		return
	}
	defer r.Body.Close()

	addedUser, err := cfg.db.CreateUser(context.Background(), database.CreateUserParams{
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Email:     user.Email,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error with Adding user:", err)
		return
	}

	//Encode Json
	respondWithJSON(w, http.StatusOK, addedUser)
}
