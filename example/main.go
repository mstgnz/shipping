package main

import (
	"log"

	"github.com/mstgnz/cargo/config"
)

func main() {

	provider, err := cargo.NewProviderByName("aras")
	if err != nil {
		return
	}

	log.Println(provider)

}
