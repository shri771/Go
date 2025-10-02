package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/shri771/Go/WebServer/internal/auth"
)

func (cfg *apiConfig) handlerWebhooks(w http.ResponseWriter, r *http.Request) {
	// Types
	type data struct {
		UserID uuid.UUID `json:"user_id"`
	}
	type Response struct {
		Event string `json:"event,omitempty"`
		Data  data   `json:"data,omitempty"`
	}

	// Validate Request
	apiKey, err := auth.GetAPIKey(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Couldn't find api key", err)
		return
	}
	if apiKey != cfg.pokaKey {
		respondWithError(w, http.StatusUnauthorized, "API key is invalid", err)
		return
	}

	// Decode incoming json
	response := Response{}
	err = json.NewDecoder(r.Body).Decode(&response)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not decode JSON", err)
		return
	}
	defer r.Body.Close()

	// Check for valid Event
	if response.Event != "user.upgraded" {
		log.Println("Not a valid Plan")
		w.WriteHeader(http.StatusNoContent)
		return
	}

	// Update User Plan
	err = cfg.db.UpdateUserByID(context.Background(), response.Data.UserID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Error with Updating Plan", err)
		return
	}

	w.WriteHeader(http.StatusNoContent)

}
