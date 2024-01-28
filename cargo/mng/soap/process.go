package soap

import (
	"github.com/mstgnz/shipping/config"
)

// CreateAbroad
// https://service.mngkargo.com.tr/musterikargosiparis/musterikargosiparis.asmx
func CreateAbroad(current cargo.Current, data cargo.ShippingData) (map[string]any, error) {
	_ = parseAbroad(data)
	return nil, nil
}

func CreateDomestic(current cargo.Current, data cargo.ShippingData) (map[string]any, error) {
	_ = parseDomestic(data)
	return nil, nil
}

func parseAbroad(data cargo.ShippingData) error {
	return nil
}

func parseDomestic(data cargo.ShippingData) error {
	return nil
}
