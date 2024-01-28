package soap

import (
	"bytes"
	"encoding/xml"
	"errors"
	"log"
	"net/http"

	"github.com/mstgnz/shipping/config"
)

// CreateAbroad global service
func CreateAbroad(current cargo.Current, data cargo.ShippingData) (map[string]any, error) {
	_, _ = parseAbroad(data)
	return nil, nil
}

// CreateDomestic local service
// TEST REQUEST: https://service.mngkargo.com.tr/tservis/musterikargosiparis.asmx?WSDL
// LIVE REQUEST: https://service.mngkargo.com.tr/musterikargosiparis/musterikargosiparis.asmx?WSDL
func CreateDomestic(current cargo.Current, data cargo.ShippingData) (map[string]any, error) {
	xmlData, err := parseDomestic(data)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(current.Endpoint.GetDevelopment(), "application/x-www-form-urlencoded", bytes.NewBuffer(xmlData))

	log.Println(resp)
	result := make(map[string]interface{})
	result["xml"] = xmlData
	return result, nil
}

func parseAbroad(data cargo.ShippingData) ([]byte, error) {
	return nil, nil
}

func parseDomestic(data cargo.ShippingData) ([]byte, error) {
	data, ok := data.(SiparisGirisiDetayliV3)
	if ok {
		xmlData, err := xml.MarshalIndent(data, "", "    ")
		if err != nil {
			return nil, err
		}
		return xmlData, nil
	}
	return nil, errors.New("data parameter not a `SiparisGirisiDetayliV3` type")
}
