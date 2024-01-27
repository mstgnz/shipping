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

func (a arasCargo) CreateCargo(data map[string]any) (map[string]any, error) {
	var err error
	var result map[string]any
	if a.GetServiceType() == cargo.SOAP {
		if a.IsDomestic() {
			result, err = soap.CreateDomestic(a.GetCurrentEndpointAndCredential(), data)
		} else {
			result, err = soap.CreateAbroad(a.GetCurrentEndpointAndCredential(), data)
		}
	}
	if a.GetServiceType() == cargo.REST {
		if a.IsDomestic() {
			result, err = rest.CreateDomestic(a.GetCurrentEndpointAndCredential(), data)
		} else {
			result, err = rest.CreateAbroad(a.GetCurrentEndpointAndCredential(), data)
		}
	}
	return result, err
}

func (a arasCargo) WhereIsTheCargo(tracking string) (map[string]any, error) {
	//TODO implement me
	panic("implement me")
}

func (a arasCargo) CancelCargo(tracking string) (map[string]any, error) {
	//TODO implement me
	panic("implement me")
}
