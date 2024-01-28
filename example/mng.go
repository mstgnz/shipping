package main

import (
	"log"

	"github.com/mstgnz/shipping/cargo/mng/soap"
	"github.com/mstgnz/shipping/config"
	"github.com/mstgnz/shipping/util"
)

func main() {

	mng, err := util.NewProviderByName("mng")
	mng.SetServiceType(cargo.SOAP) // cargo.SOAP or cargo.REST
	mng.SetMode(cargo.DEVELOPMENT) // cargo.PRODUCTION or cargo.DEVELOPMENT
	mng.SetDomestic(true)          // True if the product is domestic, False if it is abroad
	mng.AddEndpoint("test", "https://...", "https://...").SetActive(true)
	mng.AddEndpoint("live", "https://...", "https://...")
	mng.AddCredential("test user", "username", "password").SetActive(true)
	mng.AddCredential("live user", "username", "password")

	if err != nil {
		log.Println("Init Error: ", err)
	}

	s := soap.SiparisGirisiDetayliV3{}

	createCargo, err := mng.CreateCargo(s)
	if err != nil {
		log.Println("Create Error: ", err)
	}

	log.Println(createCargo)

}
