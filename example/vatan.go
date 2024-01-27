package main

import (
	"log"

	"github.com/mstgnz/shipping/util"
)

func main() {

	vatan, err := util.NewProviderByName("vatan")

	if err != nil {
		log.Println("Error: ", err)
	}

	log.Println(vatan.CreateCargo(nil))

}
