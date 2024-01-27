package main

import (
	"log"

	"github.com/mstgnz/shipping/util"
)

func main() {

	iyi, err := util.NewProviderByName("iyi")

	if err != nil {
		log.Println("Error: ", err)
	}

	log.Println(iyi.CreateCargo(nil))

}
