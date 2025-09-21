package main

import (
	"context"
	"net/http"
	"time"

	"github.com/shri771/Go/WebServer/internal/auth"
)

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

	// Check if revoked
	if tokenData.RevokedAt.Valid {
		respondWithError(w, http.StatusUnauthorized, "Session Expired", err)
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
