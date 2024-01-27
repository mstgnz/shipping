package yurtici

import (
	"github.com/mstgnz/shipping/cargo/yurtici/rest"
	"github.com/mstgnz/shipping/cargo/yurtici/soap"
	"github.com/mstgnz/shipping/config"
)

type yurticiCargo struct {
	*cargo.Cargo
}

func NewYurticiCargo() cargo.Shipper {
	return &yurticiCargo{&cargo.Cargo{}}
}

func (y yurticiCargo) CreateCargo(data map[string]any) (map[string]any, error) {
	var err error
	var result map[string]any
	if y.GetServiceType() == cargo.SOAP {
		if y.IsDomestic() {
			result, err = soap.CreateDomestic(y.GetCurrentEndpointAndCredential(), data)
		} else {
			result, err = soap.CreateAbroad(y.GetCurrentEndpointAndCredential(), data)
		}
	}
	if y.GetServiceType() == cargo.REST {
		if y.IsDomestic() {
			result, err = rest.CreateDomestic(y.GetCurrentEndpointAndCredential(), data)
		} else {
			result, err = rest.CreateAbroad(y.GetCurrentEndpointAndCredential(), data)
		}
	}
	return result, err
}

func (y yurticiCargo) WhereIsTheCargo(tracking string) (map[string]any, error) {
	//TODO implement me
	panic("implement me")
}

func (y yurticiCargo) CancelCargo(tracking string) (map[string]any, error) {
	//TODO implement me
	panic("implement me")
}
