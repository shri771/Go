package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/shri771/Go/WebServer/internal/auth"
	"github.com/shri771/Go/WebServer/internal/database"
)

func (cfg *apiConfig) handlerChirpsAdd(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Body string `json:"body"`
	}
	type returnVals struct {
		CleanedBody string `json:"cleaned_body"`
	}

	// Validate User
	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Could not retrive Auth header", err)
		return
	}

	userID, err := auth.ValidateJWT(token, cfg.secret)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Could not Validate JWT", err)
		return
	}

	// Decode Request
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err = decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}
	r.Body.Close()

	// Filter Chirps
	cleaned := validateChirps(w, params.Body)
	if cleaned == "" {
		respondWithError(w, http.StatusBadRequest, "Chirp is too long", nil)
		return
	}

	// Add Filtered Chiprs to database
	addedChipr, err := cfg.db.AddChirp(context.Background(), database.AddChirpParams{
		Body:   cleaned,
		UserID: userID,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not add Chirp to database", err)
		return
	}

	respondWithJSON(w, http.StatusCreated, Chirp{
		ID:        addedChipr.ID,
		CreatedAt: addedChipr.CreatedAt,
		UpdatedAt: addedChipr.UpdatedAt,
		Body:      addedChipr.Body,
		UserID:    addedChipr.UserID,
	})
}
