package rest

import (
	"github.com/mstgnz/shipping/config"
)

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
