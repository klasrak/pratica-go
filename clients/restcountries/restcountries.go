package restcountries

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/klasrak/apigo/models"
)

type restCountriesClient struct {
	client  *http.Client
	baseURL string
}

func (r *restCountriesClient) GetCountry(name string) ([]models.Country, error) {
	resource := fmt.Sprintf("/name/%s", name)

	res, err := r.client.Get(r.baseURL + resource)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	var country []models.Country

	err = json.Unmarshal(body, &country)

	if err != nil {
		return nil, err
	}

	return country, nil
}

func NewRestCountriesClient() *restCountriesClient {
	return &restCountriesClient{
		client:  &http.Client{},
		baseURL: "https://restcountries.com/v3.1",
	}
}
