package main

import (
	"log"

	"github.com/mstgnz/shipping/util"
)

func main() {

	mng, err := util.NewProviderByName("mng")

	if err != nil {
		log.Println("Error: ", err)
	}

	log.Println(mng.CreateCargo(nil))

}
