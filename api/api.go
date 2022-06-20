package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Api struct {
	Server http.Server
}

func (s *Api) Run() error {
	if err := s.Server.ListenAndServe(); err != nil {
		return err
	}

	return nil
}

func New() *Api {
	router := chi.NewRouter()

	initRoutes(router)

	return &Api{
		Server: http.Server{
			Addr:    ":8080",
			Handler: router,
		},
	}
}
