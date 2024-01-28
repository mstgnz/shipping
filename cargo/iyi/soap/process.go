package soap

import (
	"github.com/mstgnz/shipping/config"
)

// CreateAbroad global service
func CreateAbroad(current cargo.Current, data cargo.ShippingData) (map[string]any, error) {
	_, _ = parseAbroad(data)
	return nil, nil
}

// CreateDomestic local service
func CreateDomestic(current cargo.Current, data cargo.ShippingData) (map[string]any, error) {
	_, _ = parseDomestic(data)
	return nil, nil
}

func parseAbroad(data cargo.ShippingData) ([]byte, error) {
	return nil, nil
}

func parseDomestic(data cargo.ShippingData) ([]byte, error) {
	return nil, nil
}
