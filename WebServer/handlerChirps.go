package main

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/shri771/Go/WebServer/internal/auth"
	"github.com/shri771/Go/WebServer/internal/database"
)

type Chirp struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    uuid.UUID `json:"user_id"`
	Body      string    `json:"body"`
}

// Main Handlers
func (cfg *apiConfig) handlerChirps(w http.ResponseWriter, r *http.Request) {

	// For Get Request
	if r.Method == "GET" {
		// Get all chiprs
		chiprs, err := cfg.db.GetAllChirpAsc(context.Background())
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Could not get all Chirps in Asc", err)
			return
		}

		if len(chiprs) == 0 {
			respondWithError(w, http.StatusInternalServerError, "There are no chiprs", err)
			return
		}

		// Fill the chiprs
		jsonchirps := []Chirp{}
		for _, chirp := range chiprs {
			jsonchirps = append(jsonchirps, Chirp{
				ID:        chirp.ID,
				CreatedAt: chirp.CreatedAt,
				UpdatedAt: chirp.UpdatedAt,
				Body:      chirp.Body,
				UserID:    chirp.UserID,
			})

		}
		respondWithJSON(w, http.StatusOK, jsonchirps)
		return
	}

	// Types
	type parameters struct {
		Body string `json:"body"`
	}
	type returnVals struct {
		CleanedBody string `json:"cleaned_body"`
	}

	// Verify User
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

	// Decode Data
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err = decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

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

func (cfg *apiConfig) handlerChirpsID(w http.ResponseWriter, r *http.Request) {
	chirpIDString := r.PathValue("chirpID")
	chirpID, err := uuid.Parse(chirpIDString)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error with converting to uuid:", err)
		return
	}

	chirp, err := cfg.db.GetChiprByID(context.Background(), chirpID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Error with getting Chirp by Id:", err)
		return
	}

	respondWithJSON(w, http.StatusOK, Chirp{
		ID:        chirp.ID,
		CreatedAt: chirp.CreatedAt,
		UpdatedAt: chirp.UpdatedAt,
		Body:      chirp.Body,
		UserID:    chirp.UserID,
	})
}

// Helpers
func getCleanedBody(body string, badWords map[string]struct{}) string {
	words := strings.Split(body, " ")
	for i, word := range words {
		loweredWord := strings.ToLower(word)
		if _, ok := badWords[loweredWord]; ok {
			words[i] = "****"
		}
	}
	cleaned := strings.Join(words, " ")
	return cleaned
}

func validateChirps(w http.ResponseWriter, parms string) string {
	const maxChirpLength = 140
	if len(parms) > maxChirpLength {
		// Error
		return ""
	}

	badWords := map[string]struct{}{
		"kerfuffle": {},
		"sharbert":  {},
		"fornax":    {},
	}
	cleaned := getCleanedBody(parms, badWords)
	return cleaned

}
