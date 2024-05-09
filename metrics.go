package main

import (
	"fmt"
	"net/http"
)

type apiConfig struct {
	fileServerHits int
}

func (cfg *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Increment counter before calling the next handler
		cfg.fileServerHits++
		// Call the next handler in the chain
		next.ServeHTTP(w, r)
	})
}

func (cfg *apiConfig) hitsHandler(w http.ResponseWriter, r *http.Request) {
	htmlContent := fmt.Sprintf("<html>\n\n<body>\n    <h1>Welcome, Chirpy Admin</h1>\n    <p>Chirpy has been visited %d times!</p>\n</body>\n\n</html>\n", cfg.fileServerHits)
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, htmlContent)
}

func (cfg *apiConfig) resetHits(w http.ResponseWriter, r *http.Request) {
	cfg.fileServerHits = 0
}
