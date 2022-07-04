package models

type CountryNativeNameOpts struct {
	Official string `json:"official"`
	Common   string `json:"common"`
}

type CountryNativeName map[string]CountryNativeNameOpts

type CountryName struct {
	Common     string            `json:"common"`
	Official   string            `json:"official"`
	NativeName CountryNativeName `json:"nativeName"`
}

type CountryCurrenciesOpts struct {
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}

type CountryCurrencies map[string]CountryCurrenciesOpts

type Country struct {
	Name        CountryName       `json:"name"`
	Tld         []string          `json:"tld"`
	Cca2        string            `json:"cca2"`
	Ccn3        string            `json:"ccn3"`
	Cca3        string            `json:"cca3"`
	Cioc        string            `json:"cioc"`
	Independent bool              `json:"independent"`
	Status      string            `json:"status"`
	UnMember    bool              `json:"unMember"`
	Currencies  CountryCurrencies `json:"currencies"`
	Capital     []string          `json:"capital"`
	Region      string            `json:"region"`
	Subregion   string            `json:"subregion"`
	Flag        string            `json:"flag"`
	Population  int               `json:"population"`
	Timezones   []string          `json:"timezones"`
	Continents  []string          `json:"continents"`
}
