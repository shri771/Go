package main

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/shri771/Go/WebServer/internal/auth"
	"github.com/shri771/Go/WebServer/internal/database"
)

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

	// Create JWT
	token, err := auth.MakeJWT(login.ID, cfg.secret, 1*time.Hour)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error while Genarating JWT token", err)
	}

	// Create and Store Refreh token
	refreshToken, err := auth.MakeRefreshToken()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not genrate refrest tokens", err)
		return
	}

	refreshTokenData, err := cfg.db.CreateRefreshToken(context.Background(), database.CreateRefreshTokenParams{
		Token:     refreshToken,
		UserID:    login.ID,
		ExpiresAt: time.Now().UTC().Add(time.Hour * 24 * 60),
	})

	respondWithJSON(w, http.StatusOK, User{
		ID:           login.ID,
		CreatedAt:    login.CreatedAt,
		UpdatedAt:    login.UpdatedAt,
		Email:        login.Email,
		Token:        token,
		RefreshToken: refreshTokenData.Token,
		Member:       login.IsChipryRed,
	})

}
