package api

import (
	"github.com/go-chi/chi/v5"
)

func initRoutes(router *chi.Mux) {

	router.Route("/v1", func(r chi.Router) {
		r.Get("/hello", hello)
		r.Get("/country/{name}", getCountry)
		r.Get("/country", getCountry)
	})
}
