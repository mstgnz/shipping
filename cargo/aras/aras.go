package aras

import (
	"github.com/mstgnz/shipping/cargo/aras/rest"
	"github.com/mstgnz/shipping/cargo/aras/soap"
	"github.com/mstgnz/shipping/config"
)

type arasCargo struct {
	*cargo.Cargo
}

func NewArasCargo() cargo.Shipper {
	return &arasCargo{&cargo.Cargo{}}
}

func (a arasCargo) CreateCargo(data cargo.ShippingData) (map[string]any, error) {
	var err error
	var result map[string]any
	switch a.GetServiceType() {
	case cargo.SOAP:
		if a.IsDomestic() {
			result, err = soap.CreateDomestic(a.GetCurrentEndpointAndCredential(), data)
		} else {
			result, err = soap.CreateAbroad(a.GetCurrentEndpointAndCredential(), data)
		}
	case cargo.REST:
		if a.IsDomestic() {
			result, err = rest.CreateDomestic(a.GetCurrentEndpointAndCredential(), data)
		} else {
			result, err = rest.CreateAbroad(a.GetCurrentEndpointAndCredential(), data)
		}
	}
	return result, err
}

func (a arasCargo) WhereIsTheCargo(data cargo.ShippingData) (map[string]any, error) {
	//TODO implement me
	panic("implement me")
}

func (a arasCargo) CancelCargo(data cargo.ShippingData) (map[string]any, error) {
	//TODO implement me
	panic("implement me")
}
