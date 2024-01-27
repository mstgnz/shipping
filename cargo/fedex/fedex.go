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

func (f fedexCargo) CreateCargo(data map[string]any) (map[string]any, error) {
	var err error
	var result map[string]any
	if f.GetServiceType() == cargo.SOAP {
		if f.IsDomestic() {
			result, err = soap.CreateDomestic(f.GetCurrentEndpointAndCredential(), data)
		} else {
			result, err = soap.CreateAbroad(f.GetCurrentEndpointAndCredential(), data)
		}
	}
	if f.GetServiceType() == cargo.REST {
		if f.IsDomestic() {
			result, err = rest.CreateDomestic(f.GetCurrentEndpointAndCredential(), data)
		} else {
			result, err = rest.CreateAbroad(f.GetCurrentEndpointAndCredential(), data)
		}
	}
	return result, err
}

func (f fedexCargo) WhereIsTheCargo(tracking string) (map[string]any, error) {
	//TODO implement me
	panic("implement me")
}

func (f fedexCargo) CancelCargo(tracking string) (map[string]any, error) {
	//TODO implement me
	panic("implement me")
}
