package main

import (
	"log"

	"github.com/mstgnz/shipping/util"
)

func ExampleSurat() {

	surat, err := util.NewProviderByName("surat")

	if err != nil {
		log.Println("Error: ", err)
	}

	log.Println(surat.CreateCargo(nil))

}
