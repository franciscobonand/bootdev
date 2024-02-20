package main

import (
	"log"
	"net/http"

	"github.com/franciscobonand/bootdev/webservers/handlers"
	"github.com/go-chi/chi/v5"
)

func main() {
	cfg := &handlers.APIConfig{}

	r := chi.NewRouter()
	fsHandler := cfg.MiddlewareMetricsInc(handlers.Fileserver())
	r.Handle("/app", fsHandler)
	r.Handle("/app/*", fsHandler)

	apiRouter := chi.NewRouter()
	apiRouter.Get("/healthz", handlers.Readiness())
	apiRouter.Get("/metrics", cfg.Metrics())
	apiRouter.Get("/reset", cfg.ResetMetrics())

	r.Mount("/api", apiRouter)
	corsMux := handlers.MiddlewareCors(r)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: corsMux,
	}
	log.Printf("Server started at %s", srv.Addr)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
