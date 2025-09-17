package main

import (
	"context"
	"net/http"
)

func (cfg *apiConfig) handlerReset(w http.ResponseWriter, r *http.Request) {
	if cfg.platform != "dev" {
		respondWithError(w, http.StatusForbidden, "Only Admin can access this", nil)
	}
	cfg.fileserverHits.Store(0)

	err := cfg.db.DelUser(context.Background())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error with reseting data", nil)
	}

}
