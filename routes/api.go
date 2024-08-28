package routes

import (
    "github.com/go-chi/chi/v5"
    "net/http"
)

const(
	PathApi = "/api"
	PingApi = "/ping"
)

func ApiRoutes (r chi.Router) {
	
	r.Route(PathApi, func(r chi.Router) {
		r.Get(PingApi, func(w http.ResponseWriter, r *http.Request){
			w.Write([]byte("pong"))
		})
	})

}