package routes

import (
    "github.com/go-chi/chi/v5"
    "net/http"
)

const (
    PathAdmin = "/admin"
    PingAdmin = "/ping"
)

func AdminRoutes (r chi.Router) {
    
    r.Route(PathAdmin, func(r chi.Router) {
		r.Get(PingAdmin, func(w http.ResponseWriter, r *http.Request){
			w.Write([]byte("pong"))
		})
	})
}