package main

import (
    "database/sql"
    "github.com/rs/zerolog"
    "github.com/rs/zerolog/log"
    "net/http"
    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
)

type Server struct {
	Router *chi.Mux
	DB     *sql.DB
	Port   string
}

func main() {
    zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

    log.Info().Msg("Preparing dependencies...")
    newServer()
    
}

func newServer(){

    log.Info().Msg("Server Starting...")

    r := chi.NewRouter()
    r.Use(middleware.Logger)
    r.Get("/", func(w http.ResponseWriter, r *http.Request) {
    	w.Write([]byte("welcome"))
    })
    http.ListenAndServe(":3000", r)

}
