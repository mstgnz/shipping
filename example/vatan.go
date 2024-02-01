package main

import (
	"log"

	"github.com/mstgnz/shipping/util"
)

func ExampleVatan() {

	vatan, err := util.NewProviderByName("vatan")

	if err != nil {
		log.Println("Error: ", err)
	}

	log.Println(vatan.CreateCargo(nil))

}
