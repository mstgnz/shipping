package main

import (
	"log"

	"github.com/mstgnz/shipping/util"
)

func ExampleDhl() {

	dhl, err := util.NewProviderByName("dhl")

	if err != nil {
		log.Println("Error: ", err)
	}

	log.Println(dhl.CreateCargo(nil))

}
