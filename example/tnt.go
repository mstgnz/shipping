package main

import (
	"log"

	"github.com/mstgnz/shipping/util"
)

func main() {

	tnt, err := util.NewProviderByName("tnt")

	if err != nil {
		log.Println("Error: ", err)
	}

	log.Println(tnt.CreateCargo(nil))

}
