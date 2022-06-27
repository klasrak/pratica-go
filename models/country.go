package models

type CountryName struct {
	Common   string `json:"common"`
	Official string `json:"official"`
}

type Country struct {
	Name        CountryName `json:"name"`
	Independent bool        `json:"independent"`
}
