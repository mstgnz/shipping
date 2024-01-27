package vatan

import (
	"github.com/mstgnz/shipping/cargo/vatan/rest"
	"github.com/mstgnz/shipping/cargo/vatan/soap"
	"github.com/mstgnz/shipping/config"
)

type vatanCargo struct {
	*cargo.Cargo
}

func NewVatanCargo() cargo.Shipper {
	return &vatanCargo{&cargo.Cargo{}}
}

func (v vatanCargo) CreateCargo(data map[string]any) (map[string]any, error) {
	var err error
	var result map[string]any
	if v.GetServiceType() == cargo.SOAP {
		if v.IsDomestic() {
			result, err = soap.CreateDomestic(v.GetCurrentEndpointAndCredential(), data)
		} else {
			result, err = soap.CreateAbroad(v.GetCurrentEndpointAndCredential(), data)
		}
	}
	if v.GetServiceType() == cargo.REST {
		if v.IsDomestic() {
			result, err = rest.CreateDomestic(v.GetCurrentEndpointAndCredential(), data)
		} else {
			result, err = rest.CreateAbroad(v.GetCurrentEndpointAndCredential(), data)
		}
	}
	return result, err
}

func (v vatanCargo) WhereIsTheCargo(tracking string) (map[string]any, error) {
	//TODO implement me
	panic("implement me")
}

func (v vatanCargo) CancelCargo(tracking string) (map[string]any, error) {
	//TODO implement me
	panic("implement me")
}
