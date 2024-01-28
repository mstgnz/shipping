package ptt

import (
	"net/http"

	"github.com/mstgnz/shipping/cargo/ptt/rest"
	"github.com/mstgnz/shipping/cargo/ptt/soap"
	"github.com/mstgnz/shipping/config"
)

type pttCargo struct {
	*cargo.Cargo
}

func NewPTTCargo() cargo.Shipper {
	return &pttCargo{&cargo.Cargo{}}
}

func (p pttCargo) CreateCargo(data cargo.ShippingData) (*http.Response, error) {
	var err error
	var result *http.Response
	switch p.GetServiceType() {
	case cargo.SOAP:
		if p.IsDomestic() {
			result, err = soap.CreateDomestic(p.GetCurrentEndpointAndCredential(), data)
		} else {
			result, err = soap.CreateAbroad(p.GetCurrentEndpointAndCredential(), data)
		}
	case cargo.REST:
		if p.IsDomestic() {
			result, err = rest.CreateDomestic(p.GetCurrentEndpointAndCredential(), data)
		} else {
			result, err = rest.CreateAbroad(p.GetCurrentEndpointAndCredential(), data)
		}
	}
	return result, err
}

func (p pttCargo) WhereIsTheCargo(data cargo.ShippingData) (*http.Response, error) {
	//TODO implement me
	panic("implement me")
}

func (p pttCargo) CancelCargo(data cargo.ShippingData) (*http.Response, error) {
	//TODO implement me
	panic("implement me")
}
