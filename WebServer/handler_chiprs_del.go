package main

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerChirpsDel(w http.ResponseWriter, r *http.Request) {
	chirpIDString := r.PathValue("chirpID")
	chirpID, err := uuid.Parse(chirpIDString)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Error with converting to uuid:", err)
		return
	}
	// Validate  User
	userID, err := cfg.getUserIDJWT(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "This user does not have this chipr", err)
		return
	}

	// Check if Chipr exist and belong to that user
	dbChipr, err := cfg.db.GetChiprByID(context.Background(), chirpID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Chipr Dose not exist", err)
		return
	}
	if dbChipr.UserID != userID {
		respondWithError(w, http.StatusForbidden, "You are not authorized to do so", err)
		return
	}

	// Delete chipr
	err = cfg.db.DeleteChiprByID(context.Background(), chirpID)
	if err != nil {
		respondWithError(w, http.StatusForbidden, "Either User dose not exist or chipr", err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
