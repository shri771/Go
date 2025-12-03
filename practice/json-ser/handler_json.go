package main

import (
	"encoding/json"
	"net/http"

	respond "github.com/shri771/practice/json-ser/internal/respond"
)

func handlerAddJson(w http.ResponseWriter, r *http.Request) {

	data := make(map[string]interface{})

	// Decode json
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		respond.RespondWithError(w, http.StatusInternalServerError, "Could not decode json", err)
		return
	}

}
