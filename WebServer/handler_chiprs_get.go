package main

import (
	"context"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type Chirp struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    uuid.UUID `json:"user_id"`
	Body      string    `json:"body"`
}

// Get all chirps in Ascending Order
func (cfg *apiConfig) handlerChirpsGet(w http.ResponseWriter, r *http.Request) {
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
}

// Retrivew By chipr Id
func (cfg *apiConfig) handlerChirpsRetrivew(w http.ResponseWriter, r *http.Request) {
	chirpIDString := r.PathValue("chirpID")
	chirpID, err := uuid.Parse(chirpIDString)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Error with converting to uuid:", err)
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

// Get Chiprs by Author Id
func (cfg *apiConfig) handlerChiprsAuthor(w http.ResponseWriter, r *http.Request) {
	s := r.URL.Query().Get("author_id")
	if s == "" {
		cfg.handlerChirpsGet(w, r)
	}

	userID, err := uuid.Parse(s)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Problming with parsing url", err)
		return
	}

	// Retrive Chiprs
	requestChirps, err := cfg.db.GetChiprByUserID(context.Background(), userID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not retrive chiprs from database", err)
		return
	}

	// Curate Chiprs
	jsonChirp := []Chirp{}
	for _, chirp := range requestChirps {
		jsonChirp = append(jsonChirp, Chirp{
			ID:        chirp.ID,
			CreatedAt: chirp.CreatedAt,
			UpdatedAt: chirp.UpdatedAt,
			Body:      chirp.Body,
			UserID:    chirp.UserID,
		})
	}

	respondWithJSON(w, http.StatusOK, jsonChirp)
}
