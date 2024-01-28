package rest

import (
	"net/http"

	"github.com/mstgnz/shipping/config"
)

// CreateAbroad global service
func CreateAbroad(current cargo.Current, data cargo.ShippingData) (*http.Response, error) {
	return nil, nil
}

// CreateDomestic local service
func CreateDomestic(current cargo.Current, data cargo.ShippingData) (*http.Response, error) {
	return nil, nil
}
