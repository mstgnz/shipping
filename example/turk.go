package main

import (
	"log"

	"github.com/mstgnz/shipping/util"
)

func ExampleTurk() {

	turk, err := util.NewProviderByName("turk")

	if err != nil {
		log.Println("Error: ", err)
	}

	log.Println(turk.CreateCargo(nil))

}
