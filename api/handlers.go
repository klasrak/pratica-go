package api

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/klasrak/apigo/clients/restcountries"
)

func hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	w.Write([]byte(`{"message": "hello, world 2"}`))
}

func getCountry(w http.ResponseWriter, r *http.Request) {

	// /country/{name}
	name := chi.URLParam(r, "name")

	if name == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("content-type", "application/json")
		w.Write([]byte(`
			{
				"error": "invalid name provided"
			}
		`))
		return
	}

	country, err := restcountries.Instance.GetCountry(name)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("content-type", "application/json")
		w.Write([]byte(`
			{
				"error": "internal server error"
			}
		`))
		return
	}

	countryJson, err := json.Marshal(country)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("content-type", "application/json")
		w.Write([]byte(`
			{
				"error": "failed to marshal country"
			}
		`))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("content-type", "application/json")
	w.Write(countryJson)
}
