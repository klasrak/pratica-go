package restcountries

import "github.com/klasrak/apigo/models"

type Client interface {
	GetCountry(name string) ([]models.Country, error)
}

var (
	defaultInstance Client = NewRestCountriesClient()
	Instance               = defaultInstance
)
