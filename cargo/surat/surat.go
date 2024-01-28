package surat

import (
	"net/http"

	"github.com/mstgnz/shipping/cargo/surat/rest"
	"github.com/mstgnz/shipping/cargo/surat/soap"
	"github.com/mstgnz/shipping/config"
)

type suratCargo struct {
	*cargo.Cargo
}

func NewSuratCargo() cargo.Shipper {
	return &suratCargo{&cargo.Cargo{}}
}

func (s suratCargo) CreateCargo(data cargo.ShippingData) (*http.Response, error) {
	var err error
	var result *http.Response
	switch s.GetServiceType() {
	case cargo.SOAP:
		if s.IsDomestic() {
			result, err = soap.CreateDomestic(s.GetCurrentEndpointAndCredential(), data)
		} else {
			result, err = soap.CreateAbroad(s.GetCurrentEndpointAndCredential(), data)
		}
	case cargo.REST:
		if s.IsDomestic() {
			result, err = rest.CreateDomestic(s.GetCurrentEndpointAndCredential(), data)
		} else {
			result, err = rest.CreateAbroad(s.GetCurrentEndpointAndCredential(), data)
		}
	}
	return result, err
}

func (s suratCargo) WhereIsTheCargo(data cargo.ShippingData) (*http.Response, error) {
	//TODO implement me
	panic("implement me")
}

func (s suratCargo) CancelCargo(data cargo.ShippingData) (*http.Response, error) {
	//TODO implement me
	panic("implement me")
}
