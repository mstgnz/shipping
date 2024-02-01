package main

import (
	"log"

	"github.com/mstgnz/shipping/util"
)

func ExampleUps() {

	ups, err := util.NewProviderByName("ups")

	if err != nil {
		log.Println("Error: ", err)
	}

	log.Println(ups.CreateCargo(nil))

}
