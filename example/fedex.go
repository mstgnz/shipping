package main

import (
	"log"

	"github.com/mstgnz/shipping/util"
)

func main() {

	fedex, err := util.NewProviderByName("fedex")

	if err != nil {
		log.Println("Error: ", err)
	}

	log.Println(fedex.CreateCargo(nil))

}
