package main

import (
	"log"

	"github.com/mstgnz/shipping/cargo/mng/soap"
	"github.com/mstgnz/shipping/util"
)

func main() {

	mng, err := util.NewProviderByName("mng")

	if err != nil {
		log.Println("Error: ", err)
	}

	s := soap.SiparisGirisiDetayliV3{}

	log.Println(mng.CreateCargo(s))

}
