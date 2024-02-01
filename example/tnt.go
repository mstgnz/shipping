package main

import (
	"log"

	"github.com/mstgnz/shipping/util"
)

func ExampleTnt() {

	tnt, err := util.NewProviderByName("tnt")

	if err != nil {
		log.Println("Error: ", err)
	}

	log.Println(tnt.CreateCargo(nil))

}
