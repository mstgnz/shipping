package main

import (
	"log"

	"github.com/mstgnz/shipping/config"
)

func main() {

	provider, err := cargo.NewProviderByName("aras")
	if err != nil {
		return
	}

	log.Println(provider)

}
