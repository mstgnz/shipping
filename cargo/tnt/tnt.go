package tnt

import (
	"net/http"

	"github.com/mstgnz/shipping/cargo/tnt/rest"
	"github.com/mstgnz/shipping/cargo/tnt/soap"
	"github.com/mstgnz/shipping/config"
)

type tntCargo struct {
	*cargo.Cargo
}

func NewTNTCargo() cargo.Shipper {
	return &tntCargo{&cargo.Cargo{}}
}

func (t tntCargo) CreateCargo(data cargo.ShippingData) (*http.Response, error) {
	var err error
	var result *http.Response
	switch t.GetServiceType() {
	case cargo.SOAP:
		if t.IsDomestic() {
			result, err = soap.CreateDomestic(t.GetCurrentEndpointAndCredential(), data)
		} else {
			result, err = soap.CreateAbroad(t.GetCurrentEndpointAndCredential(), data)
		}
	case cargo.REST:
		if t.IsDomestic() {
			result, err = rest.CreateDomestic(t.GetCurrentEndpointAndCredential(), data)
		} else {
			result, err = rest.CreateAbroad(t.GetCurrentEndpointAndCredential(), data)
		}
	}
	return result, err
}

func (t tntCargo) WhereIsTheCargo(data cargo.ShippingData) (*http.Response, error) {
	//TODO implement me
	panic("implement me")
}

func (t tntCargo) CancelCargo(data cargo.ShippingData) (*http.Response, error) {
	//TODO implement me
	panic("implement me")
}
