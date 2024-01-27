package iyi

import (
	"github.com/mstgnz/shipping/cargo/iyi/rest"
	"github.com/mstgnz/shipping/cargo/iyi/soap"
	"github.com/mstgnz/shipping/config"
)

type iyiCargo struct {
	*cargo.Cargo
}

func NewIyiCargo() cargo.Shipper {
	return &iyiCargo{&cargo.Cargo{}}
}

func (i iyiCargo) CreateCargo(data map[string]any) (map[string]any, error) {
	var err error
	var result map[string]any
	if i.GetServiceType() == cargo.SOAP {
		if i.IsDomestic() {
			result, err = soap.CreateDomestic(i.GetCurrentEndpointAndCredential(), data)
		} else {
			result, err = soap.CreateAbroad(i.GetCurrentEndpointAndCredential(), data)
		}
	}
	if i.GetServiceType() == cargo.REST {
		if i.IsDomestic() {
			result, err = rest.CreateDomestic(i.GetCurrentEndpointAndCredential(), data)
		} else {
			result, err = rest.CreateAbroad(i.GetCurrentEndpointAndCredential(), data)
		}
	}
	return result, err
}

func (i iyiCargo) WhereIsTheCargo(tracking string) (map[string]any, error) {
	//TODO implement me
	panic("implement me")
}

func (i iyiCargo) CancelCargo(tracking string) (map[string]any, error) {
	//TODO implement me
	panic("implement me")
}
