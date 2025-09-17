package main

import (
	"fmt"
	"log"
	"net/http"
	"sync/atomic"

	_ "github.com/lib/pq"
)

type apiConfig struct {
	fileserverHits atomic.Int32
}

func main() {
	// Setup defaults
	const filepathRoot = '.'
	const port = "1030"
	mux := http.NewServeMux()

	apiCfg := apiConfig{
		fileserverHits: atomic.Int32{},
	}

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	// Handlers
	mux.Handle("/app/", apiCfg.middlewareMetricsInc(http.StripPrefix("/app", http.FileServer(http.Dir(filepathRoot)))))
	mux.HandleFunc("GET /api/healthz", handlerReadiness)
	mux.HandleFunc("GET /admin/metrics", http.HandlerFunc(apiCfg.handlerMetrics))
	mux.HandleFunc("POST /admin/reset", http.HandlerFunc(apiCfg.handlerReset))
	mux.HandleFunc("POST /api/validate_chirp", http.HandlerFunc(handlerChirpsValidate))

	// Logs
	log.Printf("Serving on port: %s from %v\n", port, filepathRoot)
	log.Fatal(srv.ListenAndServe()) // When there a request this calls the ServerMux's method servehttp.

}

func (cfg *apiConfig) handlerMetrics(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf(`<html>
  <body>
    <h1>Welcome, Chirpy Admin</h1>
    <p>Chirpy has been visited %d times!</p>
  </body>
</html>`, cfg.fileserverHits.Load())))
}

func (cfg *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg.fileserverHits.Add(1)

		next.ServeHTTP(w, r)
	})
}
