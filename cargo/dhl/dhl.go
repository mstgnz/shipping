package dhl

import (
	"github.com/mstgnz/shipping/cargo/dhl/rest"
	"github.com/mstgnz/shipping/cargo/dhl/soap"
	"github.com/mstgnz/shipping/config"
)

type dhlCargo struct {
	*cargo.Cargo
}

func NewDHLCargo() cargo.Shipper {
	return &dhlCargo{&cargo.Cargo{}}
}

func (d dhlCargo) CreateCargo(data map[string]any) (map[string]any, error) {
	var err error
	var result map[string]any
	if d.GetServiceType() == cargo.SOAP {
		if d.IsDomestic() {
			result, err = soap.CreateDomestic(d.GetCurrentEndpointAndCredential(), data)
		} else {
			result, err = soap.CreateAbroad(d.GetCurrentEndpointAndCredential(), data)
		}
	}
	if d.GetServiceType() == cargo.REST {
		if d.IsDomestic() {
			result, err = rest.CreateDomestic(d.GetCurrentEndpointAndCredential(), data)
		} else {
			result, err = rest.CreateAbroad(d.GetCurrentEndpointAndCredential(), data)
		}
	}
	return result, err
}

func (d dhlCargo) WhereIsTheCargo(tracking string) (map[string]any, error) {
	//TODO implement me
	panic("implement me")
}

func (d dhlCargo) CancelCargo(tracking string) (map[string]any, error) {
	//TODO implement me
	panic("implement me")
}
