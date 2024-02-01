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

	// You can use CreateCargo as a struct or a json byte or a xml byte.
	// first example use with struct
	//structExample(mng)

	// second example use with json byte
	//jsonExample(mng)

	// third example use with xml byte
	xmlExample(mng)

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

	if createResp != nil {
		log.Println("Body: ", string(createResp.Data))
	}
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

	if createResp != nil {
		log.Println("Body: ", string(createResp.Data))
	}
}

func xmlExample(mng cargo.Shipper) {
	xmlData := []byte(`<SiparisGirisiDetayliV3 xmlns="http://tempuri.org/">
      <pChIrsaliyeNo>string xml ooo</pChIrsaliyeNo>
      <pPrKiymet>string</pPrKiymet>
      <pChBarkod>string</pChBarkod>
      <pChIcerik>string</pChIcerik>
      <pGonderiHizmetSekli>string</pGonderiHizmetSekli>
      <pTeslimSekli>int</pTeslimSekli>
      <pFlAlSms>int</pFlAlSms>
      <pFlGnSms>int</pFlGnSms>
      <pKargoParcaList>string</pKargoParcaList>
      <pAliciMusteriMngNo>string</pAliciMusteriMngNo>
      <pAliciMusteriBayiNo>string</pAliciMusteriBayiNo>
      <pAliciMusteriAdi>string</pAliciMusteriAdi>
      <pChSiparisNo>string</pChSiparisNo>
      <pLuOdemeSekli>string</pLuOdemeSekli>
      <pFlAdresFarkli>string</pFlAdresFarkli>
      <pChIl>string</pChIl>
      <pChIlce>string</pChIlce>
      <pChAdres>string</pChAdres>
      <pChSemt>string</pChSemt>
      <pChMahalle>string</pChMahalle>
      <pChMeydanBulvar>string</pChMeydanBulvar>
      <pChCadde>string</pChCadde>
      <pChSokak>string</pChSokak>
      <pChTelEv>string</pChTelEv>
      <pChTelCep>string</pChTelCep>
      <pChTelIs>string</pChTelIs>
      <pChFax>string</pChFax>
      <pChEmail>string</pChEmail>
      <pChVergiDairesi>string</pChVergiDairesi>
      <pChVergiNumarasi>string</pChVergiNumarasi>
      <pFlKapidaOdeme>int</pFlKapidaOdeme>
      <pMalBedeliOdemeSekli>string</pMalBedeliOdemeSekli>
      <pPlatformKisaAdi>string</pPlatformKisaAdi>
      <pPlatformSatisKodu>string</pPlatformSatisKodu>
      <pKullaniciAdi>string</pKullaniciAdi>
      <pSifre>string</pSifre>
    </SiparisGirisiDetayliV3>`)

	createResp, err := mng.CreateCargo(xmlData)
	if err != nil {
		log.Println("Create Error: ", err)
	}

	if createResp != nil {
		log.Println("Body: ", string(createResp.Data))
	}
}
