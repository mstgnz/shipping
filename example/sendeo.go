package main

import (
	"log"

	"github.com/mstgnz/shipping/util"
)

func main() {

	sendeo, err := util.NewProviderByName("sendeo")

	if err != nil {
		log.Println("Error: ", err)
	}

	log.Println(sendeo.CreateCargo(nil))

}
