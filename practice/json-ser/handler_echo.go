package main

import (
	"net/http"

	respond "github.com/shri771/practice/json-ser/internal/respond"
)

func handlerEchoMessage(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	msg := query.Get("message")

	if msg == "" {
		respond.RespondWithError(w, http.StatusBadRequest, "Enter query", nil)
		return
	}

	respond.RespondWithJson(w, http.StatusOK, msg)

}
