package fedex

import (
	"github.com/mstgnz/shipping/cargo/fedex/rest"
	"github.com/mstgnz/shipping/cargo/fedex/soap"
	"github.com/mstgnz/shipping/config"
)

type fedexCargo struct {
	*cargo.Cargo
}

func NewFedexCargo() cargo.Shipper {
	return &fedexCargo{&cargo.Cargo{}}
}

func (f fedexCargo) CreateCargo(data cargo.ShippingData) (map[string]any, error) {
	var err error
	var result map[string]any
	switch f.GetServiceType() {
	case cargo.SOAP:
		if f.IsDomestic() {
			result, err = soap.CreateDomestic(f.GetCurrentEndpointAndCredential(), data)
		} else {
			result, err = soap.CreateAbroad(f.GetCurrentEndpointAndCredential(), data)
		}
	case cargo.REST:
		if f.IsDomestic() {
			result, err = rest.CreateDomestic(f.GetCurrentEndpointAndCredential(), data)
		} else {
			result, err = rest.CreateAbroad(f.GetCurrentEndpointAndCredential(), data)
		}
	}
	return result, err
}

func (f fedexCargo) WhereIsTheCargo(data cargo.ShippingData) (map[string]any, error) {
	//TODO implement me
	panic("implement me")
}

func (f fedexCargo) CancelCargo(data cargo.ShippingData) (map[string]any, error) {
	//TODO implement me
	panic("implement me")
}
