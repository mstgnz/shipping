package main

import (
	"log"

	"github.com/mstgnz/shipping/util"
)

func ExampleAras() {

	aras, err := util.NewProviderByName("aras")

	if err != nil {
		log.Println("Error: ", err)
	}

	log.Println(aras.CreateCargo(nil))

}
