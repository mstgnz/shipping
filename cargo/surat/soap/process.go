package soap

import (
	"github.com/mstgnz/shipping/config"
)

// CreateAbroad global service
func CreateAbroad(current cargo.Current, data cargo.ShippingData) (*cargo.Response, error) {
	return nil, nil
}

// CreateDomestic local service
func CreateDomestic(current cargo.Current, data cargo.ShippingData) (*cargo.Response, error) {
	return nil, nil
}

// WhereIsTheCargoAbroad global service
func WhereIsTheCargoAbroad(current cargo.Current, data cargo.ShippingData) (*cargo.Response, error) {
	return nil, nil
}

// WhereIsTheCargoDomestic local service
func WhereIsTheCargoDomestic(current cargo.Current, data cargo.ShippingData) (*cargo.Response, error) {
	return nil, nil
}

// CancelCargoAbroad global service
func CancelCargoAbroad(current cargo.Current, data cargo.ShippingData) (*cargo.Response, error) {
	return nil, nil
}

// CancelCargoDomestic local service
func CancelCargoDomestic(current cargo.Current, data cargo.ShippingData) (*cargo.Response, error) {
	return nil, nil
}
