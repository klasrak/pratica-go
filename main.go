package main

import (
	"github.com/klasrak/apigo/api"
)

func main() {
	server := api.New()

	if err := server.Run(); err != nil {
		panic(err)
	}
}
