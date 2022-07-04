package restcountries

import "github.com/klasrak/apigo/models"

type MockClient struct {
	GetCountryFunc func(name string) ([]models.Country, error)
}

func (c *MockClient) GetCountry(name string) ([]models.Country, error) {
	return c.GetCountryFunc(name)
}

func (c *MockClient) Use() {
	Instance = c
}

func (c *MockClient) Clean() {
	Instance = defaultInstance
}
