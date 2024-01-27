package ptt

import (
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

func (p pttCargo) CreateCargo(data map[string]any) (map[string]any, error) {
	var err error
	var result map[string]any
	if p.GetServiceType() == cargo.SOAP {
		if p.IsDomestic() {
			result, err = soap.CreateDomestic(p.GetCurrentEndpointAndCredential(), data)
		} else {
			result, err = soap.CreateAbroad(p.GetCurrentEndpointAndCredential(), data)
		}
	}
	if p.GetServiceType() == cargo.REST {
		if p.IsDomestic() {
			result, err = rest.CreateDomestic(p.GetCurrentEndpointAndCredential(), data)
		} else {
			result, err = rest.CreateAbroad(p.GetCurrentEndpointAndCredential(), data)
		}
	}
	return result, err
}

func (p pttCargo) WhereIsTheCargo(tracking string) (map[string]any, error) {
	//TODO implement me
	panic("implement me")
}

func (p pttCargo) CancelCargo(tracking string) (map[string]any, error) {
	//TODO implement me
	panic("implement me")
}
