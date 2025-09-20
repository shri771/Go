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
	})

}

// Refresh
func (cfg *apiConfig) handlerRefresh(w http.ResponseWriter, r *http.Request) {
	// Check for a token and expiry
	type params struct {
		Token string `json:"token"`
	}

	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Could retrive auth header", err)
		return
	}

	// Query database
	tokenData, err := cfg.db.GetUserFromRefreshToken(context.Background(), token)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Could retrive auth header", err)
		return
	}

	// Check for expiry
	if !time.Now().UTC().Before(tokenData.ExpiresAt) {
		respondWithError(w, http.StatusUnauthorized, "Could retrive auth header", err)
		return
	}

	tokenJWT, err := auth.MakeJWT(tokenData.UserID, cfg.secret, 1*time.Hour)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error while Genarating JWT token", err)
	}

	respondWithJSON(w, http.StatusOK, params{
		Token: tokenJWT,
	})

}

// Revoke
func (cfg *apiConfig) handlerRevoke(w http.ResponseWriter, r *http.Request) {
	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Could retrive auth header", err)
		return
	}
	if token == "" {
		respondWithError(w, http.StatusUnauthorized, "Auth header is empty", err)
		return
	}

	_, err = cfg.db.RevokeByToken(context.Background(), token)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error while setting revoked in database", err)
	}

	type empty struct {
	}
	respondWithJSON(w, http.StatusNoContent, empty{})

}
