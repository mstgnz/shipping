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
	mng.AddEndpoint("mng", "https://service.mngkargo.com.tr/musterikargosiparis/musterikargosiparis.asmx?WSDL", "https://service.mngkargo.com.tr/tservis/musterikargosiparis.asmx?WSDL").SetActive(true)
	mng.AddCredential("test user", "username", "password").SetActive(true)
	mng.AddCredential("live user", "username", "password")

	if err != nil {
		log.Println("Init Error: ", err)
	}

	order := soap.SiparisGirisiDetayliV3{}

	// this is not necessary because the already defined user information will be set in the background.
	order.WsUserName = mng.GetCurrent().Credential.GetUsername()
	order.WsPassword = mng.GetCurrent().Credential.GetPassword()

	createResp, err := mng.CreateCargo(order)
	if err != nil {
		log.Println("Create Error: ", err)
	}

	log.Println("Body: ", string(createResp.Data))

}
