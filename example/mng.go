package main

import (
	"log"

	"github.com/mstgnz/shipping/cargo/mng/soap"
	"github.com/mstgnz/shipping/config"
	"github.com/mstgnz/shipping/util"
)

func main() {

	mng, err := util.NewProviderByName("mng")
	if err != nil {
		log.Fatalln("Init Error: ", err)
	}

	mng.SetServiceType(cargo.SOAP) // cargo.SOAP or cargo.REST
	mng.SetMode(cargo.DEVELOPMENT) // cargo.PRODUCTION or cargo.DEVELOPMENT
	mng.SetDomestic(true)          // True if the product is domestic, False if it is abroad
	mng.AddEndpoint("mng", "https://service.mngkargo.com.tr/musterikargosiparis/musterikargosiparis.asmx?WSDL", "https://service.mngkargo.com.tr/tservis/musterikargosiparis.asmx?WSDL").SetActive(true)
	mng.AddCredential("test user", "username", "password").SetActive(true)
	mng.AddCredential("live user", "username", "password")

	// You can use the CreateCargo either as a struct or as a json byte.
	// first example use with struct
	structExample(mng)

	// second example use with json byte
	//jsonExample(mng)

}

func structExample(mng cargo.Shipper) {
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

func jsonExample(mng cargo.Shipper) {
	jsonData := []byte(`{
      "pChIrsaliyeNo":"",
      "pPrKiymet":"",
      "pChBarkod":"",
      "pGonderiHizmetSekli":"",
      "pTeslimSekli":0,
      "pFlAlSms":0,
      "pFlGnSms":0,
      "pKargoParcaList":"",
      "pAliciMusteriAdi":"",
      "pChSiparisNo":"",
      "pLuOdemeSekli":"",
      "pFlAdresFarkli":"",
      "pChIl":"",
      "pChIlce":"",
      "pChAdres":"",
      "pChTelCep":"",
      "pChEmail":"",
      "pMalBedeliOdemeSekli":"",
      "pFlKapidaOdeme":0,
      "pChIcerik":"",
      "pAliciMusteriMngNo":"",
      "pAliciMusteriBayiNo":"",
      "pChSemt":"",
      "pChMahalle":"",
      "pChMeydanBulvar":"",
      "pChCadde":"",
      "pChSokak":"",
      "pChFax":"",
      "pChVergiDairesi":"",
      "pChVergiNumarasi":"",
      "pPlatformKisaAdi":"",
      "pPlatformSatisKodu":"",
      "pChTelEv":"",
      "pChTelIs":""}`)

	createResp, err := mng.CreateCargo(jsonData)
	if err != nil {
		log.Println("Create Error: ", err)
	}

	log.Println("Body: ", string(createResp.Data))
}
