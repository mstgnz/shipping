package mng

import (
	"github.com/mstgnz/shipping/cargo/mng/rest"
	"github.com/mstgnz/shipping/cargo/mng/soap"
	"github.com/mstgnz/shipping/config"
)

type mngCargo struct {
	*cargo.Cargo
}

func NewMNGCargo() cargo.Shipper {
	return &mngCargo{&cargo.Cargo{}}
}

func (m mngCargo) CreateCargo(data map[string]any) (map[string]any, error) {
	var err error
	var result map[string]any
	if m.GetServiceType() == cargo.SOAP {
		if m.IsDomestic() {
			result, err = soap.CreateDomestic(m.GetCurrentEndpointAndCredential(), data)
		} else {
			result, err = soap.CreateAbroad(m.GetCurrentEndpointAndCredential(), data)
		}
	}
	if m.GetServiceType() == cargo.REST {
		if m.IsDomestic() {
			result, err = rest.CreateDomestic(m.GetCurrentEndpointAndCredential(), data)
		} else {
			result, err = rest.CreateAbroad(m.GetCurrentEndpointAndCredential(), data)
		}
	}
	return result, err
}

func (m mngCargo) WhereIsTheCargo(tracking string) (map[string]any, error) {
	//TODO implement me
	panic("implement me")
}

func (m mngCargo) CancelCargo(tracking string) (map[string]any, error) {
	//TODO implement me
	panic("implement me")
}
