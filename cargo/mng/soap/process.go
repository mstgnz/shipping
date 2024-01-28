package soap

import (
	"encoding/xml"
	"errors"

	"github.com/mstgnz/shipping/config"
)

// CreateAbroad
// https://service.mngkargo.com.tr/musterikargosiparis/musterikargosiparis.asmx
func CreateAbroad(current cargo.Current, data cargo.ShippingData) (map[string]any, error) {
	_ = parseAbroad(data)
	return nil, nil
}

func CreateDomestic(current cargo.Current, data cargo.ShippingData) (map[string]any, error) {
	xmlData, err := parseDomestic(data)
	if err != nil {
		return nil, err
	}
	result := make(map[string]interface{})
	result["xml"] = xmlData
	return result, nil
}

func parseAbroad(data cargo.ShippingData) error {
	return nil
}

func parseDomestic(data cargo.ShippingData) (string, error) {
	data, ok := data.(SiparisGirisiDetayliV3)
	if ok {
		xmlData, err := xml.MarshalIndent(data, "", "    ")
		if err != nil {
			return "", err
		}
		return string(xmlData), nil
	}
	return "", errors.New("data parameter not a `SiparisGirisiDetayliV3` type")
}
