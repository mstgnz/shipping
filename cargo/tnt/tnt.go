package tnt

import (
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

func (t tntCargo) CreateCargo(data map[string]any) (map[string]any, error) {
	var err error
	var result map[string]any
	if t.GetServiceType() == cargo.SOAP {
		if t.IsDomestic() {
			result, err = soap.CreateDomestic(t.GetCurrentEndpointAndCredential(), data)
		} else {
			result, err = soap.CreateAbroad(t.GetCurrentEndpointAndCredential(), data)
		}
	}
	if t.GetServiceType() == cargo.REST {
		if t.IsDomestic() {
			result, err = rest.CreateDomestic(t.GetCurrentEndpointAndCredential(), data)
		} else {
			result, err = rest.CreateAbroad(t.GetCurrentEndpointAndCredential(), data)
		}
	}
	return result, err
}

func (t tntCargo) WhereIsTheCargo(tracking string) (map[string]any, error) {
	//TODO implement me
	panic("implement me")
}

func (t tntCargo) CancelCargo(tracking string) (map[string]any, error) {
	//TODO implement me
	panic("implement me")
}
