package handlers

import (
	"fmt"
	"net/http"
)

func (cfg *APIConfig) Metrics() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("Hits: %d", cfg.FileserverHits)))
	})
}

func (cfg *APIConfig) ResetMetrics() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg.FileserverHits = 0
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hits reset to 0"))
	})
}

func (cfg *APIConfig) MiddlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg.FileserverHits++
		next.ServeHTTP(w, r)
	})
}
