package soap

import (
	"github.com/mstgnz/shipping/config"
)

// CreateAbroad
// https://service.mngkargo.com.tr/musterikargosiparis/musterikargosiparis.asmx
func CreateAbroad(current cargo.Current, data cargo.ShippingData) (map[string]any, error) {
	parseAbroad(data)
	return nil, nil
}

func CreateDomestic(current cargo.Current, data cargo.ShippingData) (map[string]any, error) {
	parseDomestic(data)
	return nil, nil
}

func parseAbroad(data cargo.ShippingData) {

}

func parseDomestic(data cargo.ShippingData) {

}
