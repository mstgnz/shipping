package main

import (
	"log"

	"github.com/mstgnz/shipping/config"
)

func main() {

	aras, err := cargo.NewProviderByName("aras")
	if err != nil {
		return
	}

	log.Println(aras)

}
