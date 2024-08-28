package main

import (
    "database/sql"
    "github.com/joho/godotenv"
    "github.com/rs/zerolog"
    "github.com/rs/zerolog/log"
    "github.com/yourusername/go_rest/routes"
    "net/http"
    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
    "os"
)

type Server struct {
	Router *chi.Mux
	DB     *sql.DB
	Port   string
}


func (s *Server) Initialize() {
    log.Info().Msg("Initializing server dependencies...")

    s.Router = chi.NewRouter()
    s.Router.Use(middleware.Logger)

    s.setupRoutes()
}

func (s *Server) setupRoutes() {

   routes.ApiRoutes(s.Router)
   routes.AdminRoutes(s.Router)

}

// Run starts the HTTP server
func (s *Server) Run() {
    log.Info().Msgf("Server starting on port %s", s.Port)
    if err := http.ListenAndServe(s.Port, s.Router); err != nil {
        log.Fatal().Err(err).Msg("Failed to start server")
    }
}

func main() {
    zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

    godotenv.Load()
    
    server := &Server{Port: os.Getenv("SERVICE_PORT")}
    server.Initialize()
    server.Run()
}