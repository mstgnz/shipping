# shipping - Cargo Integration

Integrating with different shipping companies on e-commerce sites can be a time-consuming and cumbersome process. Each shipping company having its separate API documentation may lead to redundant efforts in reading and implementing these documents for each of them in every project. This project simplifies the process by providing integration with all cargo companies using the Go programming language.


Features:
- **Single Integration for All Shipping Companies:** This package offers a unified solution for integrating with various shipping companies. Instead of spending time on the API documentation of different companies separately, you can access them all through a single integration.
- **Fast and Efficient:** Leveraging the performance advantages of the Go programming language, you can manage shipping processes quickly and efficiently. The primary goal of the project is to save developers time and optimize workflow.
- **Modular Structure:** The package has a modular structure, allowing you to choose and configure the desired shipping companies. Include only the necessary integrations in your project.

This project aims to make shipping integrations easier and faster, optimizing your development process.


## Cargo Companies
| Name                                        | Domestic           | Abroad             |
|---------------------------------------------|--------------------|--------------------|
| [ARAS](https://www.araskargo.com.tr/)       | :white_check_mark: | :white_check_mark: |
| [DHL](https://www.dhl.com/)                 | :white_check_mark: | :white_check_mark: |
| [FEDEX](https://www.fedex.com/)             | :white_check_mark: | :white_check_mark: |
| [MNG](https://www.mngkargo.com.tr/)         | :white_check_mark: | :white_check_mark: |
| [PTT](https://gonderitakip.ptt.gov.tr/)     | :white_check_mark: | :x:                |
| [SENDEO](https://sendeo.com.tr/)            | :white_check_mark: | :x:                |
| [SÜRAT](https://www.suratkargo.com.tr/)     | :white_check_mark: | :x:                |
| [TNT](https://www.tnt.com/)                 | :white_check_mark: | :x:                |
| [KARGO TÜRK](https://www.kargoturk.com.tr/) | :white_check_mark: | :x:                |
| [UPS](https://www.ups.com.tr/)              | :white_check_mark: | :white_check_mark: |
| [YURTİÇİ](https://www.yurticikargo.com/)    | :white_check_mark: | :white_check_mark: |

Note: Integration will be made as the documentation of the relevant companies is accessed.


## Check List
- [ ] ARAS Cargo
- [ ] DHL Cargo
- [ ] FEDEX Cargo
- [ ] MNG Cargo
- [ ] PTT Cargo
- [ ] SENDEO Cargo
- [ ] SÜRAT Cargo
- [ ] TNT Cargo
- [ ] KARGO Türk Cargo
- [ ] UPS Cargo
- [ ] YURTİÇİ Cargo


## Example:
```go
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
	jsonData := []byte(`{"pChIrsaliyeNo":"", "pPrKiymet":"", "pChBarkod":"", "pGonderiHizmetSekli":"", "pTeslimSekli":0, "pFlAlSms":0, "pFlGnSms":0, "pKargoParcaList":"", "pAliciMusteriAdi":"", "pChSiparisNo":"", "pLuOdemeSekli":"", "pFlAdresFarkli":"", "pChIl":"", "pChIlce":"", "pChAdres":"", "pChTelCep":"", "pChEmail":"", "pMalBedeliOdemeSekli":"", "pFlKapidaOdeme":0, "pChIcerik":"", "pAliciMusteriMngNo":"", "pAliciMusteriBayiNo":"", "pChSemt":"", "pChMahalle":"", "pChMeydanBulvar":"", "pChCadde":"", "pChSokak":"", "pChFax":"", "pChVergiDairesi":"", "pChVergiNumarasi":"", "pPlatformKisaAdi":"", "pPlatformSatisKodu":"", "pChTelEv":"", "pChTelIs":""}`)

	createResp, err := mng.CreateCargo(jsonData)
	if err != nil {
		log.Println("Create Error: ", err)
	}

	if createResp != nil {
		log.Println("Body: ", string(createResp.Data))
	}
}

func xmlExample(mng cargo.Shipper) {
	xmlData := []byte(`<SiparisGirisiDetayliV3xmlns="http://tempuri.org/"><pChIrsaliyeNo>stringxmlooo</pChIrsaliyeNo><pPrKiymet>string</pPrKiymet><pChBarkod>string</pChBarkod><pChIcerik>string</pChIcerik><pGonderiHizmetSekli>string</pGonderiHizmetSekli><pTeslimSekli>int</pTeslimSekli><pFlAlSms>int</pFlAlSms><pFlGnSms>int</pFlGnSms><pKargoParcaList>string</pKargoParcaList><pAliciMusteriMngNo>string</pAliciMusteriMngNo><pAliciMusteriBayiNo>string</pAliciMusteriBayiNo><pAliciMusteriAdi>string</pAliciMusteriAdi><pChSiparisNo>string</pChSiparisNo><pLuOdemeSekli>string</pLuOdemeSekli><pFlAdresFarkli>string</pFlAdresFarkli><pChIl>string</pChIl><pChIlce>string</pChIlce><pChAdres>string</pChAdres><pChSemt>string</pChSemt><pChMahalle>string</pChMahalle><pChMeydanBulvar>string</pChMeydanBulvar><pChCadde>string</pChCadde><pChSokak>string</pChSokak><pChTelEv>string</pChTelEv><pChTelCep>string</pChTelCep><pChTelIs>string</pChTelIs><pChFax>string</pChFax><pChEmail>string</pChEmail><pChVergiDairesi>string</pChVergiDairesi><pChVergiNumarasi>string</pChVergiNumarasi><pFlKapidaOdeme>int</pFlKapidaOdeme><pMalBedeliOdemeSekli>string</pMalBedeliOdemeSekli><pPlatformKisaAdi>string</pPlatformKisaAdi><pPlatformSatisKodu>string</pPlatformSatisKodu><pKullaniciAdi>string</pKullaniciAdi><pSifre>string</pSifre></SiparisGirisiDetayliV3xmlns=>`)

	createResp, err := mng.CreateCargo(xmlData)
	if err != nil {
		log.Println("Create Error: ", err)
	}

	if createResp != nil {
		log.Println("Body: ", string(createResp.Data))
	}
}
```


## Contributing
This project is open-source, and contributions are welcome. Feel free to contribute or provide feedback of any kind.


## License
This project is licensed under the APACHE License - see the [LICENSE](LICENSE) file for details.