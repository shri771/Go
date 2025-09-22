package main

import (
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/shri771/Go/WebServer/internal/auth"
)

// Chiprs
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

// Users
func (cfg *apiConfig) getUserIDJWT(header http.Header) (uuid.UUID, error) {
	// Valdidate JWT
	accessToken, err := auth.GetBearerToken(header)
	if err != nil {
		return uuid.Nil, err
	}

	userID, err := auth.ValidateJWT(accessToken, cfg.secret)
	if err != nil {
		return uuid.Nil, err
	}

	return userID, nil
}
