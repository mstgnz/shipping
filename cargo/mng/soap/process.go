package soap

import (
	"bytes"
	"encoding/xml"
	"errors"
	"net/http"

	"github.com/mstgnz/shipping/config"
)

// CreateAbroad global service
func CreateAbroad(current cargo.Current, data cargo.ShippingData) (*http.Response, error) {
	return nil, nil
}

// CreateDomestic local service
// TEST REQUEST: https://service.mngkargo.com.tr/tservis/musterikargosiparis.asmx?WSDL
// LIVE REQUEST: https://service.mngkargo.com.tr/musterikargosiparis/musterikargosiparis.asmx?WSDL
func CreateDomestic(current cargo.Current, data cargo.ShippingData) (*http.Response, error) {

	data, ok := data.(SiparisGirisiDetayliV3)
	if !ok {
		return nil, errors.New("data parameter not a `SiparisGirisiDetayliV3` type")
	}

	xmlData, err := xml.MarshalIndent(data, "", "    ")
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(current.Endpoint.GetDevelopment(), "application/x-www-form-urlencoded", bytes.NewBuffer(xmlData))
	defer func() {
		_ = resp.Body.Close()
	}()

	return resp, err
}
